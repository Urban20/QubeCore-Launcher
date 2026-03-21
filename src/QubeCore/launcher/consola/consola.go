package consola

import (
	"fmt"
	"strings"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

const (
	VERSION  = "V1.0"
	AUTOR    = "Urb@n"
	LAUNCHER = "QubeCore"
)

var Opcion1 = "lanzar version"
var Opcion2 = "ver configuracion"
var Opcion3 = "actualizar lista de versiones"
var Opcion4 = "salir"

var Pantalla = Iniciar_Pantalla()

func Menu(opciones []string) string {

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

func Imprimir_cartel(texto ...string) {
	var t string
	for _, palabra := range texto {

		t += palabra
	}

	cartel := pterm.Info
	cartel.Println(t)
}

func Iniciar_Pantalla() *pterm.AreaPrinter {

	pantalla := pterm.AreaPrinter{}
	pantalla.Fullscreen = true
	pantalla.Center = true

	p, _ := pantalla.Start()
	return p
}

func Limpiar_consola(pantalla *pterm.AreaPrinter) { // esto no funciona bien TODO
	fmt.Println(strings.Repeat("\n", 50))
	fmt.Print("\033[H")
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

func Imprimir_error(errores ...string) {
	// llamar con error.Error()
	var t string

	for _, palabra := range errores {

		t += palabra
	}

	cartel := pterm.Error
	cartel.Println(t)

}

func Imprimir_Alerta(alerta ...string) {
	var t string

	for _, palabra := range alerta {

		t += palabra
	}

	cartel := pterm.Warning
	cartel.Prefix.Text = "ADVERTENCIA"
	cartel.Println(t)

}

func Mostrar_Opciones(usuario, ruta_java, java_ram string, hilos int) {
	opciones := fmt.Sprintf("Nombre de usuario: %s\nRuta de java: %s (en la env)\nHilos en paralelo: %d\nRam de jvm: %s",
		pterm.LightMagenta(usuario), ruta_java, hilos, java_ram)

	centro := pterm.DefaultCenter
	centro.Println(opciones)
}

// resalta texto con un color celeste y letras negras
func Resaltar_texto(texto string) string {

	return pterm.FgLightCyan.Sprint(texto)

}

func Resaltar_texto_amarillo(texto string) string {

	return pterm.FgLightYellow.Sprint(texto)
}
