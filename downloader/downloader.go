package main

import (
	"fmt"
	"launcher/downloader/archivos"
	"launcher/downloader/data"
	"launcher/downloader/goruntinas"
	"os"
	"path/filepath"
)

const (
	GORUNTINAS = 1000
)

var tasks []data.Task

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Uso: mc_downloader <url-del-version-json>")
		fmt.Println("Ejemplo: mc_downloader https://piston-meta.mojang.com/v1/packages/.../1.21.10.json")
		os.Exit(1)
	}

	versionURL := os.Args[1] // cambiar luego

	//descargar version JSON

	var vj data.VersionJSON

	archivos.Obtener_Json(versionURL, &vj)
	clientPath := filepath.Join(archivos.MCDIR, "versions", vj.ID, vj.ID+".jar") // ruta versions
	// client JAR
	tasks = archivos.Cliente_JAR(tasks, vj, clientPath)

	// Guardar el version JSON localmente también
	tasks = archivos.Guarda_Json(tasks, vj, versionURL)

	tasks = archivos.Maneja_Librerias(tasks, vj)

	assetIndexPath := filepath.Join(archivos.MCDIR, "assets", "indexes", vj.AssetIndex.ID+".json")
	fmt.Printf("Fetching asset index: %s\n", vj.AssetIndex.URL)

	tasks = archivos.Maneja_Assets(tasks, vj, assetIndexPath, GORUNTINAS)

	// ── 6. Descargar todo ────────────────────────────────────────────────────
	goruntinas.RunWorkers(tasks, GORUNTINAS)

	// Armar classpath: client.jar + cada library permitida
	cp := archivos.Crear_cp(clientPath, vj)

	bat := archivos.Crear_comando(cp, vj) // vj = versionJson, cp = classpath

	batPath := "launch.bat" //
	if err := os.WriteFile(batPath, []byte(bat), 0755); err != nil {
		fmt.Println("Error generando launch.bat:", err)
		os.Exit(1)
	}
	fmt.Printf("✓ launch.bat generado. Ejecutalo para iniciar Minecraft.\n")
}
