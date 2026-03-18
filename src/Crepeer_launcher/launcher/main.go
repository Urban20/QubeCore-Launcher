package main

import (
	"downloader"
	"fmt"
	"launcher/versiones"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	exe, _         = os.Getwd()
	ruta_minecraft = filepath.Clean(filepath.Join(exe, ".minecraft"))
	ruta_versiones = filepath.Clean(filepath.Join(ruta_minecraft, "versions"))
)

func main() {

	var bytes []byte
	if !versiones.Existe_archivo(versiones.ARCHIVO_INSTANCIAS) {
		fmt.Println("json no encontrado, descargando")
		bytes = versiones.Obtener_data(versiones.VERSIONES_JSON)
		versiones.Guardar_versiones(bytes)
	} else {
		fmt.Print("\nse encontro el json")
		bytes = versiones.Leer_json(versiones.ARCHIVO_INSTANCIAS)

	}

	// impresion de versiones
	versiones_ := versiones.Listar_Versiones(bytes)
	fmt.Print("\nversiones disponibles:\n\n")

	var formato = "%d) %s\n"
	for _, version := range versiones_ {
		ruta := filepath.Clean(filepath.Join(ruta_minecraft, version.Nombre))
		//fmt.Println(ruta) ver por que no se reconocen las rutas TODO
		if versiones.Existe_archivo(ruta) {
			fmt.Printf(formato+" []", version.Indice, version.Nombre)

		} else {
			fmt.Printf(formato, version.Indice, version.Nombre)
		}

	}
	//------------------------------------------

	var eleccion int
	fmt.Print("seleccionar numero de version > ")
	fmt.Scanln(&eleccion)

	for _, v := range versiones_ {
		if v.Indice == eleccion {

			comando := downloader.Descargar_version(v.Url)
			exec.Command("java", comando...).Run()
			break
		}
	}

}
