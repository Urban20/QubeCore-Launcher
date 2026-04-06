package utilidades_test

import (
	"QbCore/utilidades"
	"strings"
	"testing"
)

func TestExtraer_version(t *testing.T) {
	tests := []struct {
		name string

		texto string
		want  string
	}{

		{name: "url", texto: "https://piston-meta.mojang.com/v1/../1.7.10.json", want: "1.7.10"},
		{name: "test", texto: "test", want: ""},
		{name: "archivo", texto: "1.21.7.json", want: "1.21.7"},
		{name: "archivo 2", texto: "1.21.json", want: "1.21"},
		{name: "version test", texto: "Paper 1.21.10", want: "1.21.10"},
		{name: "version test 2", texto: "Paper 1.21", want: "1.21"},
		{name: "url 2", texto: "https://piston-meta.mojang.com/v1/../26.1.1.json", want: "26.1.1"},
		{name: "url 3", texto: "https://piston-meta.mojang.com/v1/../27.1.json", want: "27.1"},
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
	}{ // estos tests no estan pensados para tener prefijos delante, ejemplo Paper 1.20 , etc
		{name: "verdadero 1", version: "1.20.1", want: false},
		{name: "entrada incorrecta", version: "test", want: false},
		{name: "version nueva", version: "26.1", want: false},
		{name: "version nueva 2", version: "26.1.1", want: false},
		{name: "version nueva 2", version: "27.1.1", want: false},
		{name: "version antigua", version: "1.7.2", want: true},
		{name: "version antigua 2", version: "1.7", want: true},
		{name: "version antigua 3", version: "1.5.2", want: true},
		{name: "caso borde", version: "1.8", want: true},
		{name: "limite", version: "1.8.1", want: false},
		{name: "malformado", version: "1.7.1.1", want: false},
		{name: "malformado 2", version: "26.1.1.1", want: false},
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

func TestEs_version_nueva(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		version string
		want    bool
	}{
		// estos tests no estan pensados para tener prefijos delante, ejemplo Paper 1.20 , etc
		{name: "falso 1", version: "1.20.1", want: false},
		{name: "falso 2", version: "test", want: false},
		{name: "falso 3", version: "1.20", want: false},
		{name: "verdadero 1", version: "26.1", want: true},
		{name: "verdadero 2", version: "26.1.1", want: true},
		{name: "verdadero 3", version: "27.1.1", want: true},
		{name: "malformado", version: "27.1.1.0", want: false},
		{name: "malformado 2", version: "1.1.0.1", want: false},
		{name: "verdadero 4", version: "27.1", want: true},
		{name: "verdadero 3", version: "27", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utilidades.Es_version_nueva(tt.version)

			if got != tt.want {
				t.Errorf("Es_version_nueva() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsuario_valido(t *testing.T) {
	tests := []struct {
		name string

		usuario string
		want    bool
	}{
		{name: "incorrecto 1", usuario: "Urb@n", want: false},
		{name: "incorrecto 2", usuario: "", want: false},
		{name: "incorrecto 3", usuario: strings.Repeat("test", 16), want: false},
		{name: "incorrecto 4", usuario: "test test", want: false},
		{name: "incorrecto 5", usuario: strings.Repeat("test", 16), want: false},
		{name: "incorrecto 6", usuario: "Urban@", want: false},
		{name: "correcto 1", usuario: "Urb4n_", want: true},
		{name: "correcto 2", usuario: "123", want: true},
		{name: "espacio vacio", usuario: " ", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utilidades.Usuario_valido(tt.usuario)

			if got != tt.want {
				t.Errorf("Usuario_valido() = %v, want %v", got, tt.want)
			}
		})
	}
}
