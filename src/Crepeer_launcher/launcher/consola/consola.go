package consola

import (
	"fmt"
	"launcher/versiones"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

const (
	VERSION = "V1.0"
	AUTOR   = "Urb@n"
)

func Imprimir_cartel(texto string) {
	pterm.Info.Println(texto)
}

func Mostrar_lista_Versiones(versiones_ []versiones.Versiones, ruta_versiones string, LIMITE int) {
	var contador int
	for _, version := range versiones_ {
		ruta := filepath.Join(ruta_versiones, version.Nombre)

		if versiones.Existe_archivo(ruta) {
			fmt.Printf("%d) %s   [instalada]\n", version.Indice, version.Nombre)
		} else {
			fmt.Printf("%d) %s\n", version.Indice, version.Nombre)
		}
		contador++
		if contador > LIMITE {
			fmt.Println("\nse pueden elegir otras versiones ...")
			break
		}

	}

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
