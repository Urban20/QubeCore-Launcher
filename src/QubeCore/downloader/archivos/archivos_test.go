package archivos_test

import (
	"downloader/archivos"
	"testing"
)

func TestExtraer_version(t *testing.T) {
	tests := []struct {
		name string

		archivo string
		want    string
	}{
		{name: "exito", archivo: "test 1.20.3", want: "1.20.3"},
		{name: "error", archivo: "test", want: ""},
		{name: "caso principal", archivo: "1.20.json", want: "1.20"},
		{name: "caso principal 2", archivo: "1.21.11.json", want: "1.21.11"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := archivos.Extraer_version(tt.archivo)

			if got != tt.want {
				t.Errorf("Extraer_version() = %v, want %v", got, tt.want)
			}
		})
	}
}
