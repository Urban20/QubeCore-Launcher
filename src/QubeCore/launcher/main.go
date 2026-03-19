package main

import (
	"downloader"
	"downloader/archivos"
	"fmt"
	"launcher/configuracion"
	"launcher/consola"
	"launcher/versiones"
	"os"
	"os/exec"
	"time"
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

		bytes = archivos.Descargar_Manifiest()

	} else {
		consola.Imprimir_cartel("se encontro el JSON\n")
		bytes = versiones.Leer_json(versiones.ARCHIVO_INSTANCIAS)

	}
	return bytes
}

func lanzar_versiones(bytes []byte, config configuracion.Configuracion_) {

	var version_elegida string
	versiones_ := versiones.Listar_Versiones(bytes)
	versiones.Mostrar_lista_Versiones(versiones_, versiones.Ruta_versiones, 10)

	fmt.Print("seleccionar version ➡ ")
	fmt.Scanln(&version_elegida)

	for _, v := range versiones_ {

		buscar_instancia(version_elegida, config.Usuario, config.Ruta_Java, v)
	}

}

func main() {
	configuracion.Crear_ini()
	var ejecucion bool = true

	config := configuracion.Leer_config()

	bytes := cargar_version()
	fmt.Println("Usuario iniciado como: ", config.Usuario) // TODO: decorar

	consola.Imprimir_logo()

	for ejecucion {
		consola.Limpiar_consola()
		eleccion := consola.Menu(consola.Opcion1, consola.Opcion2, consola.Opcion3)

		switch eleccion {

		case consola.Opcion1:
			lanzar_versiones(bytes, config)
		case consola.Opcion2:
			fmt.Print("\n\nno implementado todavia\n") // TODO: hacer el display de config
			fmt.Scanln()
		case consola.Opcion3:
			fmt.Print("\n\nsaliendo del launcher ...\n")
			time.Sleep(time.Second * 3)
			ejecucion = false

		}

		// esto es para el lanzamiento de versiones
		consola.Limpiar_consola()
	}
}
