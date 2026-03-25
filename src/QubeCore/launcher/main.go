package main

import (
	"QbCore/configuracion"
	"QbCore/consola"
	"QbCore/menu"
	"QbCore/versiones"
	"downloader/archivos"
	"fmt"
	"os"
)

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

func main() {

	// la configuracion se inicia al iniciar el programa en su respectivo modulo

	if err := consola.Iniciar_ANSI(); err != nil {
		err := fmt.Errorf("esta terminal no es compatible con el launcher: %w", err)
		fmt.Println(err.Error())
		fmt.Scanln()
		os.Exit(1)
	}
	fmt.Print("\033[?1049h")

	var ejecucion bool = true

	bytes := cargar_version()

	menu.Setear_opciones()

	Pantalla := consola.Iniciar_Pantalla()
	consola.Limpiar_consola(Pantalla)

	for ejecucion {

		consola.Cartel_Usuario(fmt.Sprintf("Usuario iniciado como: %s\nentrar a %s para modificarlo", consola.Color_principal.Sprint(configuracion.Config.Usuario), consola.Color_principal.Sprint(configuracion.Ruta_config)))
		consola.Imprimir_logo()

		consola.Instrucciones()

		eleccion := consola.Menu(menu.Menu_opciones)

		switch eleccion {

		case consola.Opcion1:
			menu.Lanzar_versiones(bytes)

		case consola.Opcion2:
			menu.Opcion_ver_config(Pantalla)

		case consola.Opcion3:

			menu.Opcion_actualizarVersiones(&ejecucion)

		case consola.Opcion4:
			menu.Opcion_salir(&ejecucion)

		default:
			menu.No_implementado()

		}

		consola.Limpiar_consola(Pantalla)
	}
}
