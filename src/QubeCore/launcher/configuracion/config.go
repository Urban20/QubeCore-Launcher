package configuracion

import (
	"QbCore/consola"
	"QbCore/versiones"
	"fmt"
	"strconv"

	"os"
	"time"

	"github.com/bigkevmcd/go-configparser"
)

// valores por defecto
const CONFIG = "./config.ini"

var seccion_usuario = "Usuario"
var opcion_usuario = "Nickname"
var Usuario_default = "Steve"

var seccion_ruta_java = "Ruta_Java"
var opcion_ruta_java = "Ruta"
var Ruta_dafault = "java"

var seccion_concurrencia = "Concurrencia"
var opcion_concurrencia = "Hilos"
var Hilos_default = "50"

type Configuracion_ struct {
	Usuario   string
	Ruta_Java string
	Hilos     int
}

func Crear_ini() Configuracion_ {

	if versiones.Existe_archivo(CONFIG) {
		return leer_config() // si existe la lee
	}

	ini := configparser.New() // si no existe primero la crea y despues la lee
	ini.AddSection(seccion_usuario)
	ini.Set(seccion_usuario, opcion_usuario, Usuario_default)

	ini.AddSection(seccion_ruta_java)
	ini.Set(seccion_ruta_java, opcion_ruta_java, Ruta_dafault)

	ini.AddSection(seccion_concurrencia)
	ini.Set(seccion_concurrencia, opcion_concurrencia, Hilos_default)

	ini.SaveWithDelimiter(CONFIG, "=")

	return leer_config()

}

func leer_config() Configuracion_ {
	cfg, errcfg := configparser.NewConfigParserFromFile(CONFIG)

	if errcfg != nil {
		fmt.Println("error en configuracion: ", errcfg)
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}

	conf := Configuracion_{}

	Nick, _ := cfg.Get(seccion_usuario, opcion_usuario)
	Java, _ := cfg.Get(seccion_ruta_java, opcion_ruta_java)
	Hilos_str, _ := cfg.Get(seccion_concurrencia, opcion_concurrencia)

	conf.Usuario = Nick
	conf.Ruta_Java = Java
	Hilos, errhilos := strconv.Atoi(Hilos_str)

	if errhilos != nil {
		consola.Imprimir_error("se paso un valor incorrecto al .ini")
		fmt.Scanln()
		os.Exit(1)
	}
	conf.Hilos = Hilos

	return conf

}
