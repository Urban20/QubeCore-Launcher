package menu

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

// este modulo contiene las funciones de las opciones que se llaman a main

// funciones  auxiliares

func Formatear_opciones_menu(opciones ...string) []string {

	var formateados = []string{}

	for n, op := range opciones {
		n++

		formateados = append(formateados, fmt.Sprintf("%d) %s", n, op))

	}
	return formateados

}

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
func Opcion_ver_config() {

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
}

func Opcion_actualizarVersiones(ejecucion *bool) {

	archivos.Descargar_Manifiest()
	consola.Imprimir_cartel("se debe reiniciar el launcher ...\ncerrando programa")
	time.Sleep(3 * time.Second)
	*ejecucion = false
}

func Opcion_salir(ejecucion *bool) {
	fmt.Print("\n\nsaliendo del launcher ...\n")
	time.Sleep(time.Second * 3)
	*ejecucion = false
}

func No_implementado() {
	fmt.Print("\n\n")
	consola.Imprimir_Alerta("no implementado todavia") // TODO: hacer el display de config
	fmt.Scanln()
}
