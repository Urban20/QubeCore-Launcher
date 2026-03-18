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

const LIMITE = 20 // es un limitador de impresion para no llenar la consola de versiones

func buscar_instancia(eleccion string, v versiones.Versiones) {

	if v.Nombre == eleccion {

		comando := downloader.Descargar_version(v.Url)
		cmd := exec.Command("java", comando...) // asumo que el usuario tiene java
		nul, _ := os.Open(os.DevNull)
		cmd.Stdout = nul
		cmd.Run()

	}
}

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

	var contador int // para que no muestre todas porque son un monton

	// muestra las versiones una a una
	for _, version := range versiones_ {
		//ruta := filepath.Clean(filepath.Join(ruta_minecraft, version.Nombre))
		//fmt.Println(ruta) ver por que no se reconocen las rutas TODO

		fmt.Printf(formato, version.Indice, version.Nombre)
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

			buscar_instancia(eleccion, v)
		}
	}
}
