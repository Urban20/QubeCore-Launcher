package downloader

import (
	"downloader/archivos"
	"downloader/data"
	"downloader/goruntinas"
	"launcher/versiones"
	"path"

	"path/filepath"
)

// Downloader maneja la descarga de archivos y el lanzamiento del juego

const (
	GORUNTINAS = 1000
)

var tasks []data.Task

// descarga la carpeta justo con la version y retorna el comando de lanzamiento del juego
func Descargar_version(versionURL, usuario string) []string {

	var version_json = path.Base(versionURL)
	var carpeta_version = archivos.Extraer_version(version_json)
	var ruta_target_versiones = filepath.Join(versiones.Ruta_versiones, carpeta_version)
	var ruta_target_assets = filepath.Join(versiones.Ruta_minecraft, "assets", "indexes")

	var vj data.VersionJSON

	archivos.Obtener_Json(versionURL, ruta_target_versiones, &vj)
	clientPath := filepath.Join(archivos.MCDIR, "versions", vj.ID, vj.ID+".jar") // ruta versions
	// client JAR
	tasks = archivos.Cliente_JAR(tasks, vj, clientPath)

	// Guardar el version JSON localmente también
	tasks = archivos.Guarda_Json(tasks, vj, versionURL)

	tasks = archivos.Maneja_Librerias(tasks, vj)

	assetIndexPath := filepath.Join(archivos.MCDIR, "assets", "indexes", vj.AssetIndex.ID+".json")

	tasks = archivos.Maneja_Assets(tasks, vj, assetIndexPath, ruta_target_assets, GORUNTINAS)

	// Descargar todo
	goruntinas.RunWorkers(tasks, GORUNTINAS)

	// Armar classpath: client.jar + cada library permitida
	cp := archivos.Crear_cp(clientPath, vj)

	bat := archivos.Crear_comando(usuario, cp, vj) // vj = versionJson, cp = classpath

	return bat
}
