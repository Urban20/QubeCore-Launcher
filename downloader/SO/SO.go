package so

import (
	"launcher/downloader/data"
	"runtime"
)

// identifica el sistema operativo
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
		matches := rule.OS.Name == "" || rule.OS.Name == CurrentOS()
		if matches {
			allowed = rule.Action == "allow"
		}
	}
	return allowed
}
