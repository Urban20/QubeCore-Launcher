package archivos

import (
	"encoding/json"
	"fmt"
	"launcher/downloader/data"
	"net/http"
	"os"
	"path/filepath"
)

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

func Obtener_assets(tasks []data.Task, ai data.AssetIndex, GORUNTINAS int) {

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
}
