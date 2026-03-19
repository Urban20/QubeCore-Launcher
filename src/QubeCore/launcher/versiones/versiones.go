package versiones

import (
	"encoding/json"
	"fmt"
	"io"
	"launcher/consola"
	"net/http"
	"os"
	"path/filepath"
)

const VERSIONES_JSON = "https://launchermeta.mojang.com/mc/game/version_manifest.json"
const ARCHIVO_INSTANCIAS = "./versiones.json"

type MapaVersiones struct {
	Versions []map[string]string
}

type Versiones struct { // esto contiene info de nombre (1.21.10 ejemplo , url url para descargar la version)
	Nombre string
	Url    string
	Indice int
}

var Versiones_disponibles []Versiones

// obtiene la url y devulve nil o bytes
func Obtener_data(url string) []byte {

	resp, resperr := http.Get(url)
	if resperr != nil {

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

func Existe_archivo(archivo string) bool {
	_, error := os.Stat(archivo)

	return error == nil

}

func Guardar_versiones(data []byte) {

	arch, _ := os.Create(ARCHIVO_INSTANCIAS)

	arch.Write(data)

}

// busca las versiones release y retorna una lista de estructuras
func Listar_Versiones(bytes []byte) []Versiones {

	v := MapaVersiones{}

	if json.Unmarshal(bytes, &v) != nil {
		fmt.Println("hubo un problema al listar las versiones")
		os.Exit(1)

	}

	var indice int
	for _, mapa := range v.Versions {

		version_ := mapa["id"]
		url_ := mapa["url"]
		tipo := mapa["type"]

		if tipo == "release" {

			Versiones_disponibles = append(Versiones_disponibles, Versiones{
				Nombre: version_, Url: url_, Indice: indice,
			})
			indice++

		}

	}
	return Versiones_disponibles
}

func Mostrar_lista_Versiones(versiones_ []Versiones, ruta_versiones string, LIMITE int) {
	var contador int
	for _, version := range versiones_ {
		ruta := filepath.Join(ruta_versiones, version.Nombre)

		if Existe_archivo(ruta) {
			fmt.Printf("%d) %s   [instalada]\n", version.Indice, version.Nombre)
		} else {
			fmt.Printf("%d) %s\n", version.Indice, version.Nombre)
		}
		contador++
		if contador > LIMITE {
			consola.Imprimir_cartel("\nse pueden elegir otras versiones ...")
			break
		}

	}

}
