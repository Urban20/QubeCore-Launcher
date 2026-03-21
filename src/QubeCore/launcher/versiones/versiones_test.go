package versiones_test

import (
	"QbCore/versiones"
	"testing"
)

func TestNum_version(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
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
			got := versiones.Num_version(tt.version)

			if got != tt.want {
				t.Errorf("Num_version() = %v, want %v", got, tt.want)
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
			got := versiones.Es_version_antigua(tt.version)

			if got != tt.want {
				t.Errorf("Es_version_antigua() = %v, want %v", got, tt.want)
			}
		})
	}
}
