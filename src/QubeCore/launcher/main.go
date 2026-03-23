package main

import (
	"QbCore/configuracion"
	"QbCore/consola"
	"QbCore/menu"
	"QbCore/versiones"
	"downloader/archivos"
	"fmt"
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

	var ejecucion bool = true

	bytes := cargar_version()

	menu.Setear_opciones()

	for ejecucion {
		fmt.Print("\033[?1049h")
		consola.Limpiar_consola(consola.Pantalla)
		consola.Cartel_Usuario(fmt.Sprintf("Usuario iniciado como: %s, entrar a %s para modificarlo", configuracion.Config.Usuario, configuracion.CONFIG))
		consola.Imprimir_logo()

		consola.Instrucciones()

		eleccion := consola.Menu(menu.Menu_opciones)

		switch eleccion {

		case consola.Opcion1:
			menu.Lanzar_versiones(bytes)

		case consola.Opcion2:
			menu.Opcion_ver_config()

		case consola.Opcion3:

			menu.Opcion_actualizarVersiones(&ejecucion)

		case consola.Opcion4:
			menu.Opcion_salir(&ejecucion)

		default:
			menu.No_implementado()

		}

		consola.Limpiar_consola(consola.Pantalla)
	}
}
