package versiones

import (
	"QbCore/consola"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"net/http"
	"os"
	"path/filepath"
)

var (
	Exe_archivo, _ = os.Executable()
	Exe            = filepath.Dir(Exe_archivo) //ruta del exe
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

// busca las versiones release o snapshot y retorna una lista de estructuras
func Listar_Versiones(bytes []byte, tipo string) []Versiones {

	/*
		bytes -> bytes del json de versiones
		tipo -> tipo de version: release o snapshot

	*/

	var Versiones_disponibles []Versiones

	v := MapaVersiones{}

	if json.Unmarshal(bytes, &v) != nil {
		fmt.Println("hubo un problema al listar las versiones")
		os.Exit(1)

	}

	var espacios = strings.Repeat(" ", 4)

	for _, mapa := range v.Versions {

		version_ := mapa["id"]
		url_ := mapa["url"]
		tipo_json := mapa["type"]

		if tipo_json == tipo {

			if Es_version_antigua(version_) {

				version_ = version_ + espacios + consola.Resaltar_texto_amarillo("[posibles problemas al lanzar]")

			} else if Existe_version(version_) {

				version_ = version_ + espacios + consola.Resaltar_texto("[instalada]")

			}

			Versiones_disponibles = append(Versiones_disponibles, Versiones{
				Nombre: version_, Url: url_,
			})

		}

	}
	return Versiones_disponibles
}

func Existe_archivo(archivo string) bool {
	_, error_ := os.Stat(archivo)

	return error_ == nil

}

// retorna true si la carpeta de la version existe
func Existe_version(version string) bool { // TODO: hacer un test de casos para esto

	v := filepath.Join(Ruta_versiones, version)

	return Existe_archivo(v)

}

func Menu_Versiones(versiones []Versiones) string {

	var versiones_str = []string{consola.Resaltar_texto("volver")}

	for _, version := range versiones {
		versiones_str = append(versiones_str, version.Nombre)
	}
	return consola.Menu(versiones_str)
}

func Num_version(version string) string {

	/*
		ejemplos:
		1.7 -> 7
		1.20.1 -> 20
	*/
	var valor string
	reg := regexp.MustCompile(`1\.(\d+)(?:.\d+)?`)

	matches := reg.FindStringSubmatch(version)

	if len(matches) == 2 {
		valor = matches[1]
		return valor
	}

	return valor
}

func Es_version_antigua(version string) bool {
	/* estas versiones no las cubre el laucher por falta de compatibilidad,
	las voy a dar por deprecadas (versiones menores a la 1.8 inclusive)
	(almenos por ahora)
	*/

	if version == "1.8" {
		return true
	}

	num, numerr := strconv.Atoi(Num_version(version))
	if numerr != nil {
		return false
	}

	return num <= 7

}

func Extraer_version(texto string) string {

	reg := regexp.MustCompile(`(\d+\.\d+(?:\.\d+)?)`)
	return reg.FindString(texto)

}
