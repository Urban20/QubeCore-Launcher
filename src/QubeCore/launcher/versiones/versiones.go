package versiones

import (
	"QbCore/consola"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"net/http"
	"os"
	"path/filepath"
)

var (
	Exe, _         = os.Getwd()
	Ruta_minecraft = filepath.Clean(filepath.Join(Exe, ".minecraft"))
	Ruta_versiones = filepath.Join(Ruta_minecraft, "versions")
)

const VERSIONES_JSON = "https://launchermeta.mojang.com/mc/game/version_manifest.json"
const ARCHIVO_INSTANCIAS = "./versiones.json"

type MapaVersiones struct {
	Versions []map[string]string
}

type Versiones struct { // esto contiene info de nombre (1.21.10 ejemplo , url url para descargar la version)
	Nombre string
	Url    string
}

// obtiene la url y devulve nil o bytes
func Obtener_data(url string) []byte {
	// json de versiones manifiest.json
	resp, resperr := http.Get(url)
	if resperr != nil {

		return nil
	}

	if resp.StatusCode != http.StatusOK {

		fmt.Printf("no se pudo extraer la informacion de versiones, codigo de estado: %d", resp.StatusCode) // cambiar con un mensaje de error mal lindo
		return nil
	}

	bytes, readerr := io.ReadAll(resp.Body)
	if readerr != nil {
		return nil
	}

	return bytes

}

func Leer_json(json_ string) []byte {

	data_json, _ := os.ReadFile(json_)
	return data_json

}

// descarga el json
func Guardar_versiones(data []byte) {

	arch, _ := os.Create(ARCHIVO_INSTANCIAS)

	arch.Write(data)

}

// busca las versiones release y retorna una lista de estructuras
func Listar_Versiones(bytes []byte) []Versiones {
	var Versiones_disponibles []Versiones

	v := MapaVersiones{}

	if json.Unmarshal(bytes, &v) != nil {
		fmt.Println("hubo un problema al listar las versiones")
		os.Exit(1)

	}

	for _, mapa := range v.Versions {

		version_ := mapa["id"]
		url_ := mapa["url"]
		tipo := mapa["type"]

		if tipo == "release" {

			if Existe_version(version_) {

				version_ = version_ + strings.Repeat(" ", 4) + "[instalada]"

			}

			Versiones_disponibles = append(Versiones_disponibles, Versiones{
				Nombre: version_, Url: url_,
			})

		}

	}
	return Versiones_disponibles
}

func Existe_archivo(archivo string) bool {
	_, error := os.Stat(archivo)

	return error == nil

}

func Existe_version(version string) bool {

	v := filepath.Join(Ruta_versiones, version)

	return Existe_archivo(v)

}

func Menu_Versiones(versiones []Versiones) string {

	var versiones_str = []string{"volver"}

	for _, version := range versiones {
		versiones_str = append(versiones_str, version.Nombre)
	}
	return consola.Menu(versiones_str)
}
