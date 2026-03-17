package main

import (
	"fmt"
	"launcher/versiones"
)

func main() {
	bytes := versiones.Obtener_data(versiones.VERSIONES_JSON)

	if bytes == nil {
		fmt.Println("no se pudo obtener la lista de versiones")
		return
	}
	versiones := versiones.Listar_Versiones(bytes)

	fmt.Print("\nversiones disponibles:\n\n")

	for _, version := range versiones {
		fmt.Printf("%d) %s\n", version.Indice, version.Nombre)

	}

	//downloader.Descargar_version("")

}
