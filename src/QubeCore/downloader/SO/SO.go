package so

import (
	"downloader/data"
	"fmt"
	"runtime"
)

var sistema_op = map[string]string{
	"windows": "windows",
	"darwin":  "osx",
	"linux":   "linux",
}

// esta funcion privada esta fraccionada porque se utiliza para testing
func sistemaOP(sistema string) (string, error) {

	sis := sistema_op[sistema]

	if sis == "" {
		return "", fmt.Errorf("el sistema operativo detectado no es compatible: %s", sistema)
	}
	return sis, nil

}

// identifica el sistema operativo para descargar los archivos necesarios
func SistemaOP() (string, error) {
	// ESTA ES LA FUNCION UTILIZADA PUBLICAMENTE POR EL PROGRAMA
	return sistemaOP(runtime.GOOS)

}

// verifica si la libreria esta habilitada para el sistema operativo actual
func LibraryAllowed(lib data.Library) bool {
	if len(lib.Rules) == 0 {
		return true
	}
	allowed := false
	SO, _ := SistemaOP()

	for _, rule := range lib.Rules {

		matches := rule.OS.Name == "" || rule.OS.Name == SO //pregunto si el nombre esta vacio o si coincide con el del sistema
		if matches {
			allowed = rule.Action == "allow" // busco que este en "allow" si coincide lo de arriba
		}
	}
	return allowed
}
