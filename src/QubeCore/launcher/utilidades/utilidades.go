package utilidades

import "os"

func Existe_archivo(archivo string) bool {
	_, error_ := os.Stat(archivo)

	return error_ == nil

}
