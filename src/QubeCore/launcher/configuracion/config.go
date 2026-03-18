package configuracion

import (
	"fmt"
	"os"

	"github.com/bigkevmcd/go-configparser"
)

const CONFIG = "./config.ini"

type Configuracion_ struct {
	Usuario   string
	Ruta_Java string
}

func Leer_config() Configuracion_ {
	cfg, errcfg := configparser.NewConfigParserFromFile(CONFIG)

	if errcfg != nil {
		fmt.Println("error en configuracion: ", errcfg)
		os.Exit(1)
	}

	conf := Configuracion_{}

	Nick, _ := cfg.Get("Usuario", "Nickname")
	Java, _ := cfg.Get("Java", "Ruta")

	conf.Usuario = Nick
	conf.Ruta_Java = Java

	return conf

}
