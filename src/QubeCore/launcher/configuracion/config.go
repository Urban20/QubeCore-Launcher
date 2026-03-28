package configuracion

import (
	"QbCore/consola"
	"QbCore/utilidades"
	"fmt"
	"os/exec"
	"path/filepath"
	"strconv"

	"os"
	"time"

	"github.com/bigkevmcd/go-configparser"
)

// este modulo maneja la logica de creacion y lectura de .ini (archivo de configuracion del programa)
var Exe_archivo, _ = os.Executable()
var Exe = filepath.Dir(Exe_archivo) //ruta del exe
var Config = Crear_ini()

// valores por defecto
const NOMBRE_CONFIG = "config.ini"

var (
	Ruta_config = filepath.Join(Exe, NOMBRE_CONFIG)

	// usuario
	seccion_usuario = "Usuario"
	opcion_usuario  = "Nickname"
	Usuario_default = "Steve"

	// java
	seccion_Java            = "Java"
	opcion_ruta_java        = "Ruta"
	opcion_ram_asignada     = "Ram_asignada"
	ruta_java_ejecutable, _ = exec.LookPath("java")
	Arg_default             = "2G"

	// descarga y concurrencia
	seccion_concurrencia = "Concurrencia"
	opcion_concurrencia  = "Hilos"
	Hilos_default        = "50"

	// juego

	seccion_juego      = "Minecraft"
	opcion_ruta_juego  = "Ruta"
	Ruta_juego_default = filepath.Clean(filepath.Join(Exe, ".minecraft"))
)

type Configuracion_ struct { // los valores de la config
	Usuario    string
	Ruta_Java  string
	Ram        string
	Hilos      int
	Ruta_juego string
}

func Crear_ini() Configuracion_ {

	if utilidades.Existe_archivo(Ruta_config) {
		return leer_config() // si existe la lee
	}

	ini := configparser.New() // si no existe primero la crea y despues la lee
	ini.AddSection(seccion_usuario)
	ini.Set(seccion_usuario, opcion_usuario, Usuario_default)

	// seccion java
	ini.AddSection(seccion_Java)
	ini.Set(seccion_Java, opcion_ruta_java, ruta_java_ejecutable)
	ini.Set(seccion_Java, opcion_ram_asignada, Arg_default)

	// seccion concurrencia
	ini.AddSection(seccion_concurrencia)
	ini.Set(seccion_concurrencia, opcion_concurrencia, Hilos_default)

	// seccion juego
	ini.AddSection(seccion_juego)
	ini.Set(seccion_juego, opcion_ruta_juego, Ruta_juego_default)

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

	// seteo valores de la ruta del juego
	ruta_juego, _ := cfg.Get(seccion_juego, opcion_ruta_juego)
	conf.Ruta_juego = ruta_juego

	Hilos, errhilos := strconv.Atoi(Hilos_str)

	if errhilos != nil {
		consola.Imprimir_error("se paso un valor incorrecto al .ini")
		fmt.Scanln()
		os.Exit(1)
	}
	conf.Hilos = Hilos

	return conf

}
