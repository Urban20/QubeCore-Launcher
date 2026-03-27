package archivos

import (
	"QbCore/consola"
	"QbCore/versiones"
	so "downloader/SO"
	"downloader/data"
	"downloader/fabric"
	"encoding/json"
	"errors"
	"fmt"

	"net/http"
	"os"
	"path"
	"path/filepath"
)

// maneja la logica de obtencion de json y carpeta del juego

func FetchJSON(url, ruta_target string, target interface{}) error {

	// hace un get a la url y vuelca el resultado en target

	json_arch := path.Base(url)
	archivo := filepath.Join(ruta_target, json_arch)

	if versiones.Existe_archivo(archivo) {
		// si ya existe leer de ahi

		arch, err := os.Open(archivo)

		if err != nil {
			return err
		}

		return json.NewDecoder(arch).Decode(target)
	}

	resp, err := http.Get(url) // sino descarga de ahi y la prox vez mc lo cachea automaticamnete (se supone)
	// TODO: verificar bien en algun momento
	if err != nil {
		return err
	}
	return json.NewDecoder(resp.Body).Decode(target)

}

func Obtener_Json(versionURL, ruta_target string, vj *data.VersionJSON) {

	if err := FetchJSON(versionURL, ruta_target, vj); err != nil {
		fmt.Println("Error obteniendo la version JSON:", err)
		os.Exit(1)
	}
	fmt.Printf("Version: %s | Assets: %s\n\n", vj.ID, vj.AssetIndex.ID)
}

func Crear_comando(usuario, cp, java_Ram string, vj data.VersionJSON) []string {

	var asset = filepath.Join(versiones.Ruta_minecraft, "assets")
	var dir_natives = filepath.Join(versiones.Ruta_versiones, vj.ID, "natives")

	jvm := []string{
		"-Xmx" + java_Ram,
		"-Djava.library.path=" + dir_natives,
		"-Dfile.encoding=UTF-8",
	}

	optimizacion := []string{
		"-XX:+UseG1GC",
		"-XX:+UnlockExperimentalVMOptions",
		"-XX:G1NewSizePercent=20",
		"-XX:G1ReservePercent=20",
		"-XX:MaxGCPauseMillis=50",
		"-XX:G1HeapRegionSize=32M"}

	bat := []string{"-cp", cp, "net.fabricmc.loader.impl.launch.knot.KnotClient", // TODO: en algun momento voy a tener que cambiar esto
		// el hardcodeo es fragil
		"--username", usuario,
		"--version", vj.ID,
		"--gameDir", versiones.Ruta_minecraft,
		"--assetsDir", asset,
		"--assetIndex", vj.AssetIndex.ID,
		"--uuid", "00000000-0000-0000-0000-000000000000",
		"--accessToken", "0",
		"--userType", "legacy",
		"--userProperties", "{}"}

	jvm = append(jvm, optimizacion...)
	jvm = append(jvm, bat...)

	return jvm
}

func Maneja_Assets(tasks []data.Task, vj data.VersionJSON, assetIndexPath, ruta_target string, GORUNTINAS int) []data.Task {

	var ai data.AssetIndex // assetindex
	if err := FetchJSON(vj.AssetIndex.URL, ruta_target, &ai); err != nil {
		consola.Imprimir_error(fmt.Errorf("Error lanzando el indice de assets: %w", err).Error())
		consola.Imprimir_Alerta("CUIDADO!! esto puede hacer que el juego crashee")

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
		dest := filepath.Join(versiones.Ruta_minecraft, "assets", "objects", prefix, hash)
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
			DestPath: filepath.Join(versiones.Ruta_minecraft, "libraries", filepath.FromSlash(a.Path)),
			SHA1:     a.SHA1,
			Label:    a.Path,
		})
	}
	fmt.Printf("Libraries: %d a descargar, %d salteadas (otro OS)\n", len(tasks)-2, skipped)
	return tasks
}

func Guarda_Json(tasks []data.Task, vj data.VersionJSON, versionURL string) []data.Task {

	versionJSONPath := filepath.Join(versiones.Ruta_minecraft, "versions", vj.ID, vj.ID+".json")
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
		cp += string(filepath.ListSeparator) + filepath.Join(versiones.Ruta_minecraft, "libraries", filepath.FromSlash(a.Path))
	}

	for _, lib := range fabric.Iniciar_sistema_fabric() {
		cp += string(filepath.Separator) + lib
	}
	return cp
}

func Descargar_Manifiest() ([]byte, error) {

	// descarga, guarda y retorna bytes

	bytes, dataerr := versiones.Obtener_data(versiones.VERSIONES_JSON)

	if dataerr != nil {

		return []byte{}, dataerr
	}

	if bytes == nil || len(bytes) == 0 {
		return []byte{}, errors.New("se retorno un numero de bytes vacio")
	}

	versiones.Guardar_versiones(bytes)

	return bytes, nil

}
