package main

import (
	"downloader"
	"fmt"
	"launcher/configuracion"
	"launcher/versiones"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	exe, _         = os.Getwd()
	ruta_minecraft = filepath.Clean(filepath.Join(exe, ".minecraft"))
	ruta_versiones = filepath.Join(ruta_minecraft, "versions")
)

const LIMITE = 20 // es un limitador de impresion para no llenar la consola de versiones

func buscar_instancia(eleccion, usuario, ruta_java string, v versiones.Versiones) {

	if v.Nombre == eleccion {

		comando := downloader.Descargar_version(v.Url, usuario)
		cmd := exec.Command(ruta_java, comando...) // asumo que el usuario tiene java
		nul, _ := os.Open(os.DevNull)
		cmd.Stdout = nul
		cmderr := cmd.Run()

		if cmderr != nil {
			fmt.Println(cmderr)
		}

	}
}

func cargar_version() []byte {
	var bytes []byte
	if !versiones.Existe_archivo(versiones.ARCHIVO_INSTANCIAS) {
		// si el json de versiones no existe obtiene el json de internet
		fmt.Println("json no encontrado, descargando")

		bytes = versiones.Obtener_data(versiones.VERSIONES_JSON)

		versiones.Guardar_versiones(bytes)

	} else {
		fmt.Print("\nse encontro el json")
		bytes = versiones.Leer_json(versiones.ARCHIVO_INSTANCIAS)

	}
	return bytes
}

func main() {

	config := configuracion.Leer_config()

	bytes := cargar_version()

	// impresion de versiones
	versiones_ := versiones.Listar_Versiones(bytes)
	fmt.Print("\nversiones disponibles:\n\n")

	var contador int // para que no muestre todas porque son un monton

	// muestra las versiones una a una
	for _, version := range versiones_ {
		ruta := filepath.Join(ruta_versiones, version.Nombre)

		if versiones.Existe_archivo(ruta) {
			fmt.Printf("%d) %s   [instalada]\n", version.Indice, version.Nombre)
		} else {
			fmt.Printf("%d) %s\n", version.Indice, version.Nombre)
		}
		contador++
		if contador > LIMITE {
			fmt.Println("\nse pueden elegir otras versiones ...")
			break
		}

	}
	//------------------------------------------

	for {
		var eleccion string
		fmt.Print("seleccionar version > ")
		_, scanerr := fmt.Scanln(&eleccion)
		if scanerr != nil {

			continue
		}

		for _, v := range versiones_ {

			buscar_instancia(eleccion, config.Usuario, config.Ruta_Java, v)
		}
	}
}
