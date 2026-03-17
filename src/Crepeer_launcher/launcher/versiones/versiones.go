package versiones

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const VERSIONES_JSON = "https://launchermeta.mojang.com/mc/game/version_manifest.json"

type MapaVersiones struct {
	Versions []map[string]string
}

type Versiones struct { // esto contiene info de nombre (1.21.10 ejemplo , url url para descargar la version)
	Nombre string
	Url    string
	Indice int
}

var versiones_disponibles []Versiones

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

			versiones_disponibles = append(versiones_disponibles, Versiones{
				Nombre: version_, Url: url_, Indice: indice,
			})
			indice++

		}

	}
	return versiones_disponibles
}
