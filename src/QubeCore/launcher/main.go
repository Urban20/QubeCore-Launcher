package main

import (
	"QbCore/configuracion"
	"QbCore/consola"
	"QbCore/versiones"

	"downloader"
	"downloader/archivos"
	"fmt"
	"os"
	"os/exec"
	"time"
)

var config = configuracion.Crear_ini()

const LIMITE = 20 // es un limitador de impresion para no llenar la consola de versiones
var menu_opciones = []string{
	consola.Opcion1,
	consola.Opcion2,
	consola.Opcion3,
	"buscar actualizaciones de versiones"}

func buscar_instancia(interrumpido *bool, eleccion, ruta_java string, v versiones.Versiones) {
	var comando []string

	if v.Nombre == "volver" {

		*interrumpido = true

	} else if v.Nombre == eleccion {

		comando = downloader.Descargar_version(v.Url, config.Usuario, config.Hilos)
		cmd := exec.Command(ruta_java, comando...) // asumo que el usuario tiene java
		nul, _ := os.Open(os.DevNull)
		cmd.Stdout = nul
		defer nul.Close()
		consola.Imprimir_cartel("iniciando instancia...")
		cmderr := cmd.Run()

		if cmderr != nil {
			consola.Imprimir_error("hubo un problema al lanzar la version: ", cmderr.Error())
			fmt.Print("\n\n")
			consola.Imprimir_Alerta("el problema puede ser causado por una version de java incompatible")
			fmt.Scanln()
			*interrumpido = true
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

func lanzar_versiones(bytes []byte) {

	var interrumpido = false

	var version_elegida string
	versiones_ := versiones.Listar_Versiones(bytes)

	version_elegida = versiones.Menu_Versiones(versiones_)

	for _, v := range versiones_ {

		if interrumpido {
			break
		}

		buscar_instancia(&interrumpido, version_elegida, config.Ruta_Java, v)
	}

}

func main() {

	var ejecucion bool = true

	bytes := cargar_version()

	for ejecucion {
		fmt.Print("\033[?1049h")
		consola.Limpiar_consola(consola.Pantalla)
		consola.Cartel_Usuario(fmt.Sprintf("Usuario iniciado como: %s, entrar a %s para modificarlo", config.Usuario, configuracion.CONFIG))
		consola.Imprimir_logo()
		eleccion := consola.Menu(menu_opciones)

		switch eleccion {

		case consola.Opcion1:
			lanzar_versiones(bytes)

		case consola.Opcion3:
			fmt.Print("\n\nsaliendo del launcher ...\n")
			time.Sleep(time.Second * 3)
			ejecucion = false

		default:
			fmt.Print("\n\n")
			consola.Imprimir_Alerta("no implementado todavia") // TODO: hacer el display de config
			fmt.Scanln()

		}

		// esto es para el lanzamiento de versiones
		consola.Limpiar_consola(consola.Pantalla)
	}
}
