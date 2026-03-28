package utilidades_test

import (
	"QbCore/utilidades"
	"testing"
)

func TestExtraer_version(t *testing.T) {
	tests := []struct {
		name string

		texto string
		want  string
	}{

		{name: "url", texto: "https://piston-meta.mojang.com/v1/packages/ed5d8789ed29872ea2ef1c348302b0c55e3f3468/1.7.10.json", want: "1.7.10"},
		{name: "test", texto: "test", want: ""},
		{name: "archivo", texto: "1.21.7.json", want: "1.21.7"},
		{name: "archivo 2", texto: "1.21.json", want: "1.21"},
		{name: "version test", texto: "Paper 1.21.10", want: "1.21.10"},
		{name: "version test 2", texto: "Paper 1.21", want: "1.21"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utilidades.Extraer_version(tt.texto)

			if got != tt.want {
				t.Errorf("Extraer_version() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEs_version_antigua(t *testing.T) {
	tests := []struct {
		name string

		version string
		want    bool
	}{
		{name: "verdadero 1", version: "1.20.1", want: false},
		{name: "entrada incorrecta", version: "test", want: false},
		{name: "version nueva", version: "26.1", want: false},
		{name: "version antigua", version: "1.7.2", want: true},
		{name: "version antigua 2", version: "1.7", want: true},
		{name: "version antigua 3", version: "1.5.2", want: true},
		{name: "caso borde", version: "1.8", want: true},
		{name: "limite", version: "1.8.1", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utilidades.Es_version_antigua(tt.version)

			if got != tt.want {
				t.Errorf("Es_version_antigua() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNum_version(t *testing.T) {
	tests := []struct {
		name string

		version string
		want    string
	}{
		{name: "exito", version: "1.20.1", want: "20"},
		{name: "exito 2", version: "Paper 1.21.1", want: "21"},
		{name: "exito 3", version: "paper 1.20", want: "20"},
		{name: "exito 4", version: "1.21", want: "21"},
		{name: "fallo", version: "test", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utilidades.Num_version(tt.version)

			if got != tt.want {
				t.Errorf("Num_version() = %v, want %v", got, tt.want)
			}
		})
	}
}
