package consola

import (
	"fmt"
	"strings"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

const (
	VERSION  = "V1.1"
	AUTOR    = "Urb@n"
	LAUNCHER = "QubeCore"
)

var Opcion1 = "lanzar version"
var Opcion2 = "ver configuracion"
var Opcion3 = "actualizar lista de versiones"
var Opcion4 = "salir"

var Color_principal = pterm.NewRGB(131, 184, 39)
var Pantalla = Iniciar_Pantalla()

func Menu(opciones []string) string {

	seleccion := pterm.DefaultInteractiveSelect
	seleccion.TextStyle = &pterm.Style{pterm.BgGreen, pterm.FgBlack}
	seleccion.MaxHeight = 10
	seleccion.DefaultText = "SELECCIONAR opcion"
	seleccion.Selector = "➡ "
	seleccion.SelectorStyle = &pterm.Style{pterm.FgWhite}
	seleccion.FilterInputPlaceholder = "[TIPEAR opcion]"
	op, _ := seleccion.WithOptions(opciones).Show()
	return op

}

func Imprimir_cartel(texto ...string) {
	var t string
	for _, palabra := range texto {

		t += palabra
	}

	cartel := pterm.Info
	cartel.Prefix.Style = &pterm.Style{pterm.BgGreen}
	cartel.MessageStyle = &pterm.Style{pterm.FgGreen}
	cartel.Println(t)
}

func Iniciar_Pantalla() *pterm.AreaPrinter {

	pantalla := pterm.AreaPrinter{}
	pantalla.Fullscreen = true
	pantalla.Center = true

	p, _ := pantalla.Start()
	return p
}

func Limpiar_consola(pantalla *pterm.AreaPrinter) {
	fmt.Println(strings.Repeat("\n", 50))
	fmt.Print("\033[H")
	pantalla.Clear()

}

func Imprimir_logo() {
	banner := fmt.Sprintf("Launcher CLI para Minecraft Java\nVersion: %s\nEscrito por: %s", VERSION, AUTOR)
	//Limpiar_consola()
	logo, _ := pterm.DefaultBigText.WithLetters(putils.LettersFromStringWithRGB(LAUNCHER, Color_principal)).Srender()
	pterm.DefaultCenter.Println(logo)

	pterm.DefaultCenter.WithCenterEachLineSeparately().Println(pterm.Gray(banner))
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

func Impresion_centro(str string) {

	centro := pterm.DefaultCenter
	centro.Println(str)
}

// resalta texto con un color celeste y letras negras
func Resaltar_texto(texto string) string {

	return Color_principal.Sprint(texto)

}

func Resaltar_texto_amarillo(texto string) string {

	return pterm.FgLightYellow.Sprint(texto)
}

func Instrucciones() {
	fmt.Print(strings.Repeat("\n", 4))
	texto := "Teclas ←↑→↓, ENTER para confirmar"

	instrucciones := pterm.Info
	instrucciones.Prefix.Text = "INSTRUCCIONES"
	instrucciones.Prefix.Style = &pterm.Style{pterm.BgLightGreen, pterm.FgBlack}
	instrucciones.MessageStyle = &pterm.Style{pterm.FgDarkGray}
	instrucciones.Println(texto + "\n\n")

}

func Crear_barra(total int, titulo string) *pterm.ProgressbarPrinter {

	barra := pterm.DefaultProgressbar
	barra.BarStyle = &pterm.Style{pterm.BgWhite}
	barra.TitleStyle = &pterm.Style{pterm.FgBlack, pterm.BgLightWhite}
	p, _ := barra.WithTotal(total).WithTitle(titulo).Start()

	return p

}
