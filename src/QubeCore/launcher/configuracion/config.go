package configuracion

import (
	"fmt"
	"launcher/versiones"
	"os"
	"time"

	"github.com/bigkevmcd/go-configparser"
)

const CONFIG = "./config.ini"

var seccion_usuario = "Usuario"
var opcion_usuario = "Nickname"

var seccion_ruta_java = "Ruta_Java"
var opcion_ruta_java = "Ruta"

var Usuario_default = "Steve"
var Ruta_dafault = "java"

type Configuracion_ struct {
	Usuario   string
	Ruta_Java string
}

func Crear_ini() {

	if versiones.Existe_archivo(CONFIG) {
		return
	}

	ini := configparser.New()
	ini.AddSection(seccion_usuario)
	ini.AddSection(seccion_ruta_java)

	ini.Set(seccion_usuario, opcion_usuario, Usuario_default)
	ini.Set(seccion_ruta_java, opcion_ruta_java, Ruta_dafault)

	ini.SaveWithDelimiter(CONFIG, "=")

}

func Leer_config() Configuracion_ {
	cfg, errcfg := configparser.NewConfigParserFromFile(CONFIG)

	if errcfg != nil {
		fmt.Println("error en configuracion: ", errcfg)
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}

	conf := Configuracion_{}

	Nick, _ := cfg.Get(seccion_usuario, opcion_usuario)
	Java, _ := cfg.Get(seccion_ruta_java, opcion_ruta_java)

	conf.Usuario = Nick
	conf.Ruta_Java = Java

	return conf

}
