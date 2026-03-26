package main

import (
	"QbCore/configuracion"
	"QbCore/consola"
	"QbCore/menu"
	"QbCore/versiones"
	"downloader/archivos"
	"errors"
	"fmt"
	"os"
)

// invoca Descargar_manifiest manejando sus errores
func manejar_error_manifiest() []byte {

	bytesmanifiest, manifiesterr := archivos.Descargar_Manifiest()

	if manifiesterr != nil {
		consola.Imprimir_error("hubo un problema al descargar el manifiest: ", manifiesterr.Error())
		fmt.Scanln()
		return []byte{}
	}

	return bytesmanifiest

}

func cargar_version() ([]byte, error) {
	var bytes []byte
	if !versiones.Existe_archivo(versiones.ARCHIVO_INSTANCIAS) {
		// si el json de versiones no existe obtiene el json de internet
		consola.Imprimir_cartel("json no encontrado, descargando\n")

		bytes = manejar_error_manifiest()

	} else {
		consola.Imprimir_cartel("se encontro el JSON\n")
		bytes = versiones.Leer_json(versiones.ARCHIVO_INSTANCIAS)

	}

	if len(bytes) == 0 {
		return []byte{}, errors.New("se devolvio un numero de bytes vacios")
	}

	return bytes, nil
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

	bytes, versionerr := cargar_version()

	if versionerr != nil {
		consola.Imprimir_error("error al cargar las versiones: ", versionerr.Error())
		fmt.Scanln()
		os.Exit(1)

	}

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

			if vererr := menu.Lanzar_versiones(bytes); vererr != nil {
				consola.Imprimir_error("hubo un problema al lanzar las versiones: ", vererr.Error())
				fmt.Scanln()
			}

		case consola.Opcion2:
			menu.Opcion_ver_config(Pantalla)

		case consola.Opcion3:

			bytes = manejar_error_manifiest()

		case consola.Opcion4:
			menu.Opcion_salir(&ejecucion)

		default:
			menu.No_implementado()

		}

		consola.Limpiar_consola(Pantalla)
	}
}
