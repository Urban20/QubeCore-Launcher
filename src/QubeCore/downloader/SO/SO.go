package so

import (
	"downloader/data"
	"runtime"
)

// identifica el sistema operativo para descargar los archivos necesarios
func CurrentOS() string {
	switch runtime.GOOS {
	case "windows":
		return "windows"
	case "darwin":
		return "osx"
	default:
		return "linux"
	}
}

// verifica si la libreria esta habilitada para el sistema operativo actual
func LibraryAllowed(lib data.Library) bool {
	if len(lib.Rules) == 0 {
		return true
	}
	allowed := false
	for _, rule := range lib.Rules {
		matches := rule.OS.Name == "" || rule.OS.Name == CurrentOS() //pregunto si el nombre esta vacio o si coincide con el del sistema
		if matches {
			allowed = rule.Action == "allow" // busco que este en "allow" si coincide lo de arriba
		}
	}
	return allowed
}
