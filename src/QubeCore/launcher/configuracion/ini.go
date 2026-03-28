package configuracion

import (
	"QbCore/consola"
	"QbCore/utilidades"
	"fmt"
	"path/filepath"
	"strconv"

	"os"
	"time"

	"github.com/bigkevmcd/go-configparser"
)

func normalizar_ruta_juego(ruta_juego string) string {

	ruta_juego = filepath.Clean(ruta_juego)
	carpeta_juego := ".minecraft"

	if filepath.Base(ruta_juego) == carpeta_juego {
		return ruta_juego
	}

	return filepath.Join(ruta_juego, carpeta_juego)

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

	conf.Ruta_juego = normalizar_ruta_juego(ruta_juego)

	Hilos, errhilos := strconv.Atoi(Hilos_str)

	if errhilos != nil {
		consola.Imprimir_error("se paso un valor incorrecto al .ini")
		fmt.Scanln()
		os.Exit(1)
	}
	conf.Hilos = Hilos

	return conf

}
