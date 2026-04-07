package utilidades

import (
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// modulo que tiene pequeñas funciones utiles para casos concretos
var reg_versiones = regexp.MustCompile(`(\d+\.\d+(?:\.\d+)?)`)
var reg_versiones_viejas = regexp.MustCompile(`1\.(\d+)(?:\.\d+)?`)

func Existe_archivo(archivo string) bool {
	_, error_ := os.Stat(archivo)

	return error_ == nil

}

func Extraer_version(texto string) string {

	return reg_versiones.FindString(texto)

}

func Num_version(version string) string {

	/*
		ejemplos:
		1.7 -> 7
		1.20.1 -> 20
	*/
	var valor string

	matches := reg_versiones_viejas.FindStringSubmatch(version)

	if len(matches) == 2 {
		valor = matches[1]
		return valor
	}

	return valor
}

func Es_version_nueva(version string) bool {

	if !slices.Contains([]int{2, 3}, len(strings.Split(version, "."))) {
		return false
	}

	if !reg_versiones.Match([]byte(version)) {
		return false
	}

	num_version_moderno, converr := strconv.Atoi(strings.Split(version, ".")[0])

	if converr != nil {
		return false
	}

	if num_version_moderno < 26 {
		return false
	}

	return true
}

func Es_version_antigua(version string) bool {
	/* estas versiones no las cubre el laucher por falta de compatibilidad,
	las voy a dar por deprecadas (versiones menores a la 1.8 inclusive)
	(almenos por ahora)
	*/

	if !slices.Contains([]int{2, 3}, len(strings.Split(version, "."))) {
		return false
	}

	if Es_version_nueva(version) {
		return false
	}

	if version == "1.8" {
		return true
	}

	num, _ := strconv.Atoi(Num_version(version))

	return num <= 7

}

func Usuario_valido(usuario string) bool {

	/*
	   para que un usuario sea valido debe tener:

	   * de 3 a 16 caracteres
	   * sin caracteres especiales salvo "_"


	*/

	reg := regexp.MustCompile(`^[a-zA-Z0-9_]{3,16}$`)

	return reg.Match([]byte(usuario))

}
