package menu

import (
	"QbCore/configuracion"
	"QbCore/consola"
	"QbCore/versiones"
	"downloader"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/pterm/pterm"
)

// este modulo contiene las funciones de las opciones que se llaman a main

var Archivo_CMD = filepath.Join(versiones.Exe, "log-cmd.log") // desvio el comando del stdout
var Archivo_Stederr_CMD = filepath.Join(versiones.Exe, "log-cmd-error.log")

// funciones  auxiliares

func Formatear_opciones_menu(opciones ...string) []string {

	var formateados = []string{}

	for n, op := range opciones {
		n++

		formateados = append(formateados, fmt.Sprintf("%d) %s", n, op))

	}
	return formateados

}

func ejecutar_comando(ruta_java string, comando []string) error {
	cmd := exec.Command(ruta_java, comando...) // asumo que el usuario tiene java
	out, _ := os.Create(Archivo_CMD)
	stederr, _ := os.Create(Archivo_Stederr_CMD)
	cmd.Stdout = out
	cmd.Stderr = stederr

	consola.Imprimir_cartel("iniciando instancia...")
	cmderr := cmd.Run()
	out.Close()
	stederr.Close()
	return cmderr

}

func buscar_instancia(interrumpido *bool, eleccion, ruta_java string, v versiones.Versiones) {
	var comando []string

	if v.Nombre == "volver" {

		*interrumpido = true

	} else if v.Nombre == eleccion {

		comando = downloader.Descargar_version(v.Url, configuracion.Config.Usuario, configuracion.Config.Ram, configuracion.Config.Hilos)
		cmderr := ejecutar_comando(ruta_java, comando)

		if cmderr != nil {

			consola.Imprimir_error("hubo un problema al lanzar la version: ", cmderr.Error())
			fmt.Print("\n\n")
			consola.Imprimir_Alerta("el problema puede ser causado por una version de java incompatible")
			fmt.Scanln()
			*interrumpido = true
		}

	}
}

func Lanzar_versiones(bytes []byte) {

	var interrumpido = false

	var version_elegida string

	tipo := consola.Menu([]string{"release", "snapshot"})

	versiones_ := versiones.Listar_Versiones(bytes, tipo)

	version_elegida = versiones.Menu_Versiones(versiones_)

	for _, v := range versiones_ {

		if interrumpido {
			break
		}

		buscar_instancia(&interrumpido, version_elegida, configuracion.Config.Ruta_Java, v)
	}

}

var Menu_opciones = Formatear_opciones_menu(
	consola.Opcion1,
	consola.Opcion2,
	consola.Opcion3,
	consola.Opcion4)

func Setear_opciones() {
	consola.Opcion1 = Menu_opciones[0]
	consola.Opcion2 = Menu_opciones[1]
	consola.Opcion3 = Menu_opciones[2]
	consola.Opcion4 = Menu_opciones[3]

}

// opciones
func Opcion_ver_config(pantalla *pterm.AreaPrinter) {

	consola.Limpiar_consola(pantalla)
	consola.Mostrar_Opciones(
		configuracion.Config.Usuario,
		configuracion.Config.Ruta_Java,
		configuracion.Config.Ram,
		configuracion.Config.Hilos,
	)
	fmt.Print("\n\n")
	consola.Imprimir_cartel("ENTER para volver")
	fmt.Scanln()
}

/*
func Opcion_actualizarVersiones(ejecucion *bool) {

	bytes_manifiest := archivos.Descargar_Manifiest()

}}*/

func Opcion_salir(ejecucion *bool) {
	fmt.Print("\n\n")
	consola.Imprimir_cartel("saliendo del launcher ...")
	time.Sleep(time.Second * 3)
	*ejecucion = false
}

func No_implementado() {
	fmt.Print("\n\n")
	consola.Imprimir_Alerta("no implementado todavia")
	fmt.Scanln()
}
