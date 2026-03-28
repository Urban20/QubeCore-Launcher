package versiones

import (
	"QbCore/configuracion"
	"QbCore/consola"
	"QbCore/utilidades"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"net/http"
	"os"
	"path/filepath"
)

var (
	Ruta_minecraft = configuracion.Config.Ruta_juego
	Ruta_versiones = filepath.Join(Ruta_minecraft, "versions")
	Ruta_libraries = filepath.Join(Ruta_minecraft, "libraries")
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
func Obtener_data(url string) ([]byte, error) {
	// json de versiones manifiest.json
	resp, resperr := http.Get(url)
	if resperr != nil {

		return []byte{}, resperr
	}

	if resp.StatusCode != http.StatusOK {

		return []byte{}, fmt.Errorf("no se pudo extraer la informacion de versiones, codigo de estado: %d", resp.StatusCode) // cambiar con un mensaje de error mal lindo

	}

	bytes, readerr := io.ReadAll(resp.Body)
	if readerr != nil {
		return []byte{}, readerr
	}

	return bytes, nil

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

// busca las versiones release o snapshot y retorna una lista de estructuras
func Listar_Versiones(bytes []byte, tipo string) ([]Versiones, error) {

	/*
		bytes -> bytes del json de versiones
		tipo -> tipo de version: release o snapshot

	*/

	var Versiones_disponibles []Versiones

	v := MapaVersiones{}

	if err := json.Unmarshal(bytes, &v); err != nil {
		return []Versiones{}, errors.New("hubo un problema al listar las versiones")
	}

	var espacios = strings.Repeat(" ", 4)

	for _, mapa := range v.Versions {

		version_ := mapa["id"]
		url_ := mapa["url"]
		tipo_json := mapa["type"]

		if tipo_json == tipo {

			if utilidades.Es_version_antigua(version_) {

				version_ = version_ + espacios + consola.Resaltar_texto_amarillo("[posibles problemas al lanzar]")

			} else if Existe_version(version_) {

				version_ = version_ + espacios + consola.Resaltar_texto("[instalada]")

			}

			Versiones_disponibles = append(Versiones_disponibles, Versiones{
				Nombre: version_, Url: url_,
			})

		}

	}
	return Versiones_disponibles, nil
}

// retorna true si la carpeta de la version existe
func Existe_version(version string) bool {

	v := filepath.Join(Ruta_versiones, version)

	return utilidades.Existe_archivo(v)

}

func Menu_Versiones(versiones []Versiones) string {

	var versiones_str = []string{consola.Resaltar_texto("volver")}

	for _, version := range versiones {
		versiones_str = append(versiones_str, version.Nombre)
	}
	return consola.Menu(versiones_str)
}
