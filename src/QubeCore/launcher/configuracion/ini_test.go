package configuracion

import (
	"path/filepath"
	"testing"
)

func Test_normalizar_ruta_juego(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		ruta_juego string
		want       string
	}{
		{name: "test 1", ruta_juego: "docs/files/.minecraft", want: filepath.Clean("docs/files/.minecraft")},
		{name: "test 2", ruta_juego: "docs/files", want: filepath.Clean("docs/files/.minecraft")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalizar_ruta_juego(tt.ruta_juego)

			if got != tt.want {
				t.Errorf("normalizar_ruta_juego() = %v, want %v", got, tt.want)
			}
		})
	}
}
