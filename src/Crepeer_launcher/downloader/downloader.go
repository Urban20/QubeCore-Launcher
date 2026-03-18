package downloader

import (
	"downloader/archivos"
	"downloader/data"
	"downloader/goruntinas"
	"fmt"

	"path/filepath"
)

const (
	GORUNTINAS = 1000
)

var tasks []data.Task

// descarga la carpeta justo con la version y retorna el comando de lanzamiento del juego
func Descargar_version(versionURL string) []string {

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

	// Descargar todo
	goruntinas.RunWorkers(tasks, GORUNTINAS)

	// Armar classpath: client.jar + cada library permitida
	cp := archivos.Crear_cp(clientPath, vj)

	bat := archivos.Crear_comando(cp, vj) // vj = versionJson, cp = classpath

	return bat
}
