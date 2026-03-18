package consola

import (
	"fmt"
	"os/exec"
	"runtime"

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

	seleccion, _ := pterm.DefaultInteractiveSelect.WithOptions(opciones).Show()
	return seleccion

}

func Imprimir_cartel(texto string) {
	pterm.Info.Println(texto)
}

func Limpiar_consola() {

	var comando string

	switch runtime.GOOS {
	case "windows":
		comando = "cls"
	default:
		comando = "clear"
	}

	exec.Command(comando).Run()

}

func Imprimir_logo() {
	banner := fmt.Sprintf("Launcher CLI\nVersion: %s\nAutor: %s", VERSION, AUTOR)
	Limpiar_consola()
	logo, _ := pterm.DefaultBigText.WithLetters(putils.LettersFromString("Crepeer")).Srender()
	pterm.DefaultCenter.Println(logo)
	pterm.DefaultCenter.WithCenterEachLineSeparately().Println(banner)
}
