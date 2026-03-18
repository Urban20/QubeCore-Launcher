package archivos

import (
	so "downloader/SO"
	"downloader/data"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

// maneja la logica de obtencion de json y carpeta del juego

const MCDIR = "./.minecraft" // ruta del .minecraft (la misma que el programa)

func FetchJSON(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

func Obtener_Json(versionURL string, vj *data.VersionJSON) {

	if err := FetchJSON(versionURL, vj); err != nil {
		fmt.Println("Error obteniendo la version JSON:", err)
		os.Exit(1)
	}
	fmt.Printf("Version: %s | Assets: %s\n\n", vj.ID, vj.AssetIndex.ID)
}

func Crear_comando(usuario, cp string, vj data.VersionJSON) []string {

	var asset = filepath.Join(MCDIR, "assets")

	bat := []string{"-cp", cp, vj.MainClass,
		"--username", usuario,
		"--version", vj.ID,
		"--gameDir", MCDIR,
		"--assetsDir", asset,
		"--assetIndex", vj.AssetIndex.ID,
		"--uuid", "00000000-0000-0000-0000-000000000000",
		"--accessToken", "0",
		"--userType", "legacy"}

	return bat
}

func Maneja_Assets(tasks []data.Task, vj data.VersionJSON, assetIndexPath string, GORUNTINAS int) []data.Task {

	var ai data.AssetIndex // assetindex
	if err := FetchJSON(vj.AssetIndex.URL, &ai); err != nil {
		fmt.Println("Error fetching asset index:", err)
		os.Exit(1)
	}

	tasks = append(tasks, data.Task{
		URL:      vj.AssetIndex.URL,
		DestPath: assetIndexPath,
		SHA1:     vj.AssetIndex.SHA1,
		Label:    "assets/indexes/" + vj.AssetIndex.ID + ".json",
	})

	for _, obj := range ai.Objects {
		hash := obj.Hash
		prefix := hash[:2]
		url := fmt.Sprintf("https://resources.download.minecraft.net/%s/%s", prefix, hash)
		dest := filepath.Join(MCDIR, "assets", "objects", prefix, hash)
		tasks = append(tasks, data.Task{

			URL:      url,
			DestPath: dest,
			SHA1:     hash,
			Label:    hash,
		})
	}
	fmt.Printf("Assets: %d archivos\n", len(ai.Objects))
	fmt.Printf("\nTotal tareas: %d | Workers: %d\n\n", len(tasks), GORUNTINAS)

	return tasks

}

func Maneja_Librerias(tasks []data.Task, vj data.VersionJSON) []data.Task {
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
			DestPath: filepath.Join(MCDIR, "libraries", filepath.FromSlash(a.Path)),
			SHA1:     a.SHA1,
			Label:    a.Path,
		})
	}
	fmt.Printf("Libraries: %d a descargar, %d salteadas (otro OS)\n", len(tasks)-2, skipped)
	return tasks
}

func Guarda_Json(tasks []data.Task, vj data.VersionJSON, versionURL string) []data.Task {

	versionJSONPath := filepath.Join(MCDIR, "versions", vj.ID, vj.ID+".json")
	tasks = append(tasks, data.Task{
		URL:      versionURL,
		DestPath: versionJSONPath,
		Label:    vj.ID + ".json",
	})
	return tasks
}

func Cliente_JAR(tasks []data.Task, vj data.VersionJSON, clientPath string) []data.Task {

	tasks = append(tasks, data.Task{
		URL:      vj.Downloads.Client.URL,
		DestPath: clientPath,
		SHA1:     vj.Downloads.Client.SHA1,
		Label:    "client.jar",
	})
	return tasks
}

func Crear_cp(clientPath string, vj data.VersionJSON) string { // nota: cp = classpath

	cp := clientPath
	for _, lib := range vj.Libraries {
		if !so.LibraryAllowed(lib) {
			continue
		}
		a := lib.Downloads.Artifact
		if a.URL == "" {
			continue
		}
		cp += ";" + filepath.Join(MCDIR, "libraries", filepath.FromSlash(a.Path))
	}
	return cp
}
