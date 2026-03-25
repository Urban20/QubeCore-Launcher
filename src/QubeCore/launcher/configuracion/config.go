package configuracion

import (
	"QbCore/consola"
	"QbCore/versiones"
	"fmt"
	"os/exec"
	"path/filepath"
	"strconv"

	"os"
	"time"

	"github.com/bigkevmcd/go-configparser"
)

// este modulo maneja la logica de creacion y lectura de .ini (archivo de configuracion del programa)

var Config = Crear_ini()

// valores por defecto
const NOMBRE_CONFIG = "config.ini"

var Ruta_config = filepath.Join(versiones.Exe, NOMBRE_CONFIG)

// usuario
var seccion_usuario = "Usuario"
var opcion_usuario = "Nickname"
var Usuario_default = "Steve"

// java
var seccion_Java = "Java"
var opcion_ruta_java = "Ruta"
var opcion_ram_asignada = "Ram_asignada"
var ruta_java_ejecutable, _ = exec.LookPath("java")
var Arg_default = "2G"

// descarga y concurrencia
var seccion_concurrencia = "Concurrencia"
var opcion_concurrencia = "Hilos"
var Hilos_default = "50"

type Configuracion_ struct { // los valores de la config
	Usuario   string
	Ruta_Java string
	Ram       string
	Hilos     int
}

func Crear_ini() Configuracion_ {

	if versiones.Existe_archivo(Ruta_config) {
		return leer_config() // si existe la lee
	}

	ini := configparser.New() // si no existe primero la crea y despues la lee
	ini.AddSection(seccion_usuario)
	ini.Set(seccion_usuario, opcion_usuario, Usuario_default)

	ini.AddSection(seccion_Java)
	ini.Set(seccion_Java, opcion_ruta_java, ruta_java_ejecutable)
	ini.Set(seccion_Java, opcion_ram_asignada, Arg_default)

	ini.AddSection(seccion_concurrencia)
	ini.Set(seccion_concurrencia, opcion_concurrencia, Hilos_default)

	ini.SaveWithDelimiter(Ruta_config, "=")

	return leer_config()

}

func leer_config() Configuracion_ {
	cfg, errcfg := configparser.NewConfigParserFromFile(Ruta_config)

	if errcfg != nil {
		fmt.Println("error en configuracion: ", errcfg)
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}

	conf := Configuracion_{}

	// obtener usuario
	Nick, _ := cfg.Get(seccion_usuario, opcion_usuario)

	// obtener java
	ruta_Java, _ := cfg.Get(seccion_Java, opcion_ruta_java)
	Ram, _ := cfg.Get(seccion_Java, opcion_ram_asignada)

	//obtener concurrencia
	Hilos_str, _ := cfg.Get(seccion_concurrencia, opcion_concurrencia)

	//seteo de configuracion en la struct
	conf.Usuario = Nick
	conf.Ruta_Java = ruta_Java
	conf.Ram = Ram
	Hilos, errhilos := strconv.Atoi(Hilos_str)

	if errhilos != nil {
		consola.Imprimir_error("se paso un valor incorrecto al .ini")
		fmt.Scanln()
		os.Exit(1)
	}
	conf.Hilos = Hilos

	return conf

}
