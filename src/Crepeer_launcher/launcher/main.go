package main

import (
	"downloader"
	"fmt"
	"launcher/configuracion"
	"launcher/consola"
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
		consola.Imprimir_cartel("json no encontrado, descargando\n")

		bytes = versiones.Obtener_data(versiones.VERSIONES_JSON)

		versiones.Guardar_versiones(bytes)

	} else {
		consola.Imprimir_cartel("se encontro el JSON\n")
		bytes = versiones.Leer_json(versiones.ARCHIVO_INSTANCIAS)

	}
	return bytes
}

func main() {

	config := configuracion.Leer_config()
	consola.Imprimir_logo()

	bytes := cargar_version()

	// impresion de versiones
	versiones_ := versiones.Listar_Versiones(bytes)
	consola.Mostrar_lista_Versiones(versiones_, ruta_versiones, 10)
	// muestra las versiones una a una

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
