package consola

import (
	"fmt"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

const (
	VERSION = "V1.0"
	AUTOR   = "Urb@n"
)

var Opcion1 = "1) Lanzar version"
var Opcion2 = "2) Ver configuracion"
var Opcion3 = "3) Salir"

func Menu(opciones ...string) string {

	seleccion := pterm.DefaultInteractiveSelect
	seleccion.DefaultText = "SELECCIONAR opcion"
	seleccion.Selector = ">> "
	seleccion.FilterInputPlaceholder = "TIPEAR opcion"
	op, _ := seleccion.WithOptions(opciones).Show()
	return op

}

func Imprimir_cartel(texto string) {
	pterm.Info.Println(texto)
}

func Limpiar_consola() { // esto no funciona bien TODO

	area, _ := pterm.DefaultArea.Start()
	area.Clear()

}

func Imprimir_logo() {
	banner := fmt.Sprintf("Launcher CLI para Minecraft Java\nVersion: %s\nAutor: %s", VERSION, AUTOR)
	//Limpiar_consola()
	logo, _ := pterm.DefaultBigText.WithLetters(putils.LettersFromString("Crepeer")).Srender()
	pterm.DefaultCenter.Println(logo)
	pterm.DefaultCenter.WithCenterEachLineSeparately().Println(banner)
}
