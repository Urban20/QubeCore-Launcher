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

var menu_opciones = formatear_opciones_menu(
	consola.Opcion1,
	consola.Opcion2,
	consola.Opcion3,
	consola.Opcion4)

func buscar_instancia(interrumpido *bool, eleccion, ruta_java string, v versiones.Versiones) {
	var comando []string

	if v.Nombre == "volver" {

		*interrumpido = true

	} else if v.Nombre == eleccion {

		comando = downloader.Descargar_version(v.Url, configuracion.Config.Usuario, configuracion.Config.Ram, configuracion.Config.Hilos)
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

		buscar_instancia(&interrumpido, version_elegida, configuracion.Config.Ruta_Java, v)
	}

}

func formatear_opciones_menu(opciones ...string) []string {

	var formateados = []string{}

	for n, op := range opciones {
		n++

		formateados = append(formateados, fmt.Sprintf("%d) %s", n, op))

	}
	return formateados

}

func setear_opciones() {
	consola.Opcion1 = menu_opciones[0]
	consola.Opcion2 = menu_opciones[1]
	consola.Opcion3 = menu_opciones[2]
	consola.Opcion4 = menu_opciones[3]

}

func main() {

	var ejecucion bool = true

	bytes := cargar_version()

	for ejecucion {
		fmt.Print("\033[?1049h")
		consola.Limpiar_consola(consola.Pantalla)
		consola.Cartel_Usuario(fmt.Sprintf("Usuario iniciado como: %s, entrar a %s para modificarlo", configuracion.Config.Usuario, configuracion.CONFIG))
		consola.Imprimir_logo()

		setear_opciones()

		eleccion := consola.Menu(menu_opciones)

		switch eleccion {

		case consola.Opcion1:
			lanzar_versiones(bytes)

		case consola.Opcion2:
			consola.Limpiar_consola(consola.Pantalla)
			consola.Mostrar_Opciones(
				configuracion.Config.Usuario,
				configuracion.Config.Ruta_Java,
				configuracion.Config.Ram,
				configuracion.Config.Hilos,
			)
			fmt.Print("\n\n")
			consola.Imprimir_cartel("ENTER para volver")
			fmt.Scanln()
		case consola.Opcion3:

			archivos.Descargar_Manifiest()
			consola.Imprimir_cartel("se debe reiniciar el launcher ...\ncerrando programa")
			time.Sleep(3 * time.Second)
			ejecucion = false

		case consola.Opcion4:
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
