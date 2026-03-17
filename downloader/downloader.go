package main

import (
	"fmt"
	so "launcher/downloader/SO"
	"launcher/downloader/archivos"
	"launcher/downloader/data"
	"launcher/downloader/red"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
)

const (
	GORUNTINAS = 32
)

// ─── Helpers ─────────────────────────────────────────────────────────────────

// ─── Worker pool ──────────────────────────────────────────────────────────────

// descarga los archivos con concurrencia para agilizar
func runWorkers(tasks []data.Task, workers int) {
	ch := make(chan data.Task, len(tasks))
	for _, t := range tasks {
		ch <- t
	}
	close(ch)

	var (
		wg     sync.WaitGroup
		done   atomic.Int64
		total  = int64(len(tasks))
		mu     sync.Mutex
		errors []string
	)

	for i := 0; i < workers; i++ {
		wg.Add(1)

		go func() {

			defer wg.Done()
			for task := range ch {
				err := red.DownloadFile(task.URL, task.DestPath, task.SHA1)

				n := done.Add(1)
				mu.Lock()

				if err != nil {
					errors = append(errors, fmt.Sprintf("FALLO [%s]: %v", task.Label, err))
					fmt.Printf("\r[%d/%d] ✗ %s\n", n, total, task.Label)
					continue
				}

				fmt.Printf("\r[%d/%d] ✓ %-60s", n, total, task.Label)

				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println()

	if len(errors) > 0 {
		fmt.Printf("\n%d error(s):\n", len(errors))
		for _, e := range errors {
			fmt.Println(" ", e)
		}
	}
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Uso: mc_downloader <url-del-version-json>")
		fmt.Println("Ejemplo: mc_downloader https://piston-meta.mojang.com/v1/packages/.../1.21.10.json")
		os.Exit(1)
	}

	versionURL := os.Args[1] // cambiar luego

	// ── 1. Descargar version JSON ─────────────────────────────────────────────

	var vj data.VersionJSON

	archivos.Obtener_Json(versionURL, &vj)

	var tasks []data.Task

	// ── 2. Client JAR ─────────────────────────────────────────────────────────
	clientPath := filepath.Join(archivos.MCDIR, "versions", vj.ID, vj.ID+".jar") // ruta versions
	tasks = append(tasks, data.Task{
		URL:      vj.Downloads.Client.URL,
		DestPath: clientPath,
		SHA1:     vj.Downloads.Client.SHA1,
		Label:    "client.jar",
	})

	// Guardar el version JSON localmente también
	versionJSONPath := filepath.Join(archivos.MCDIR, "versions", vj.ID, vj.ID+".json")
	tasks = append(tasks, data.Task{
		URL:      versionURL,
		DestPath: versionJSONPath,
		Label:    vj.ID + ".json",
	})

	// ── 3. Libraries ──────────────────────────────────────────────────────────
	skipped := 0
	for _, lib := range vj.Libraries {
		if !so.LibraryAllowed(lib) {
			skipped++
			continue
		}
		a := lib.Downloads.Artifact
		if a.URL == "" {
			continue
		}
		tasks = append(tasks, data.Task{
			URL:      a.URL,
			DestPath: filepath.Join(archivos.MCDIR, "libraries", filepath.FromSlash(a.Path)),
			SHA1:     a.SHA1,
			Label:    a.Path,
		})
	}
	fmt.Printf("Libraries: %d a descargar, %d salteadas (otro OS)\n", len(tasks)-2, skipped)

	assetIndexPath := filepath.Join(archivos.MCDIR, "assets", "indexes", vj.AssetIndex.ID+".json")
	fmt.Printf("Fetching asset index: %s\n", vj.AssetIndex.URL)

	var ai data.AssetIndex
	if err := archivos.FetchJSON(vj.AssetIndex.URL, &ai); err != nil {
		fmt.Println("Error fetching asset index:", err)
		os.Exit(1)
	}

	tasks = append(tasks, data.Task{
		URL:      vj.AssetIndex.URL,
		DestPath: assetIndexPath,
		SHA1:     vj.AssetIndex.SHA1,
		Label:    "assets/indexes/" + vj.AssetIndex.ID + ".json",
	})

	archivos.Obtener_assets(tasks, ai, GORUNTINAS) // tareas , assets index, concurrencia

	// ── 6. Descargar todo ────────────────────────────────────────────────────
	runWorkers(tasks, GORUNTINAS)

	// ── 7. Generar launch.bat con classpath completo ──────────────────────────
	fmt.Println("\n✓ Descarga completa. Generando launch.bat...")

	// Armar classpath: client.jar + cada library permitida
	cp := clientPath
	for _, lib := range vj.Libraries {
		if !so.LibraryAllowed(lib) {
			continue
		}
		a := lib.Downloads.Artifact
		if a.URL == "" {
			continue
		}
		cp += ";" + filepath.Join(archivos.MCDIR, "libraries", filepath.FromSlash(a.Path))
	}

	bat := fmt.Sprintf(`@echo off
java -cp "%s" %s ^
  --username TuNombre ^
  --version %s ^
  --gameDir %s ^
  --assetsDir %s ^
  --assetIndex %s ^
  --uuid 00000000-0000-0000-0000-000000000000 ^
  --accessToken 0 ^
  --userType legacy
`,
		cp,
		vj.MainClass,
		vj.ID,
		archivos.MCDIR,
		filepath.Join(archivos.MCDIR, "assets"),
		vj.AssetIndex.ID,
	)

	batPath := "launch.bat"
	if err := os.WriteFile(batPath, []byte(bat), 0755); err != nil {
		fmt.Println("Error generando launch.bat:", err)
		os.Exit(1)
	}
	fmt.Printf("✓ launch.bat generado. Ejecutalo para iniciar Minecraft.\n")
}
