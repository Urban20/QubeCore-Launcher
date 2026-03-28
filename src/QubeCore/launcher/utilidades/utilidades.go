package utilidades

import (
	"os"
	"regexp"
	"strconv"
)

// modulo que tiene pequeñas funciones utiles para casos concretos

func Existe_archivo(archivo string) bool {
	_, error_ := os.Stat(archivo)

	return error_ == nil

}

func Extraer_version(texto string) string {

	reg := regexp.MustCompile(`(\d+\.\d+(?:\.\d+)?)`)
	return reg.FindString(texto)

}

func Num_version(version string) string {

	/*
		ejemplos:
		1.7 -> 7
		1.20.1 -> 20
	*/
	var valor string
	reg := regexp.MustCompile(`1\.(\d+)(?:.\d+)?`)

	matches := reg.FindStringSubmatch(version)

	if len(matches) == 2 {
		valor = matches[1]
		return valor
	}

	return valor
}

func Es_version_antigua(version string) bool {
	/* estas versiones no las cubre el laucher por falta de compatibilidad,
	las voy a dar por deprecadas (versiones menores a la 1.8 inclusive)
	(almenos por ahora)
	*/

	if version == "1.8" {
		return true
	}

	num, numerr := strconv.Atoi(Num_version(version))
	if numerr != nil {
		return false
	}

	return num <= 7

}
