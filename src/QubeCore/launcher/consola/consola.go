package consola

import (
	"fmt"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

const (
	VERSION  = "V1.0"
	AUTOR    = "Urb@n"
	LAUNCHER = "QubeCore"
)

var Opcion1 = "1) Lanzar version"
var Opcion2 = "2) Ver configuracion"
var Opcion3 = "3) Salir"

var Pantalla = Iniciar_Pantalla()

func Menu(opciones ...string) string {

	seleccion := pterm.DefaultInteractiveSelect
	seleccion.TextStyle = &pterm.Style{pterm.BgLightCyan, pterm.FgBlack}
	seleccion.DefaultText = "SELECCIONAR opcion"
	seleccion.Selector = "➡ "
	seleccion.SelectorStyle = &pterm.Style{pterm.FgWhite}
	seleccion.FilterInputPlaceholder = "[TIPEAR opcion]"
	op, _ := seleccion.WithOptions(opciones).Show()
	return op

}

func Imprimir_cartel_2(texto, nombre_cartel string) {
	// lo mismo que imprimir cartel pero adaptado a otro cartel distinto de info

	cartel := pterm.Info
	cartel.Prefix.Text = nombre_cartel
	cartel.Println(texto)

}

func Imprimir_cartel(texto string) {
	cartel := pterm.Info
	cartel.Println(texto)
}

func Iniciar_Pantalla() *pterm.AreaPrinter {

	pantalla := pterm.AreaPrinter{}
	pantalla.Fullscreen = true
	pantalla.Center = true

	p, _ := pantalla.Start()
	return p
}

func Limpiar_consola(pantalla *pterm.AreaPrinter) { // esto no funciona bien TODO

	fmt.Print("\033[H")
	fmt.Print("\033[?1049h")
	pantalla.Clear()

}

func Imprimir_logo() {
	banner := fmt.Sprintf("Launcher CLI para Minecraft Java\nVersion: %s\nAutor: %s", VERSION, AUTOR)
	//Limpiar_consola()
	logo, _ := pterm.DefaultBigText.WithLetters(putils.LettersFromString(LAUNCHER)).Srender()
	pterm.DefaultCenter.Println(logo)
	pterm.DefaultCenter.WithCenterEachLineSeparately().Println(banner)
}

func Cartel_Usuario(usuario string) {

	cartel := pterm.Info
	cartel.Prefix.Text = "USUARIO"
	cartel.Prefix.Style = &pterm.Style{pterm.BgLightMagenta, pterm.FgBlack}
	cartel.MessageStyle = &pterm.Style{pterm.FgLightWhite}
	cartel.Println(usuario)
	fmt.Print("\n\n\n")

}
