package configuracion

import (
	"os"
	"os/exec"
	"path/filepath"
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
	Ruta_juego_default = filepath.Clean(Exe)
)

type Configuracion_ struct { // los valores de la config
	Usuario    string
	Ruta_Java  string
	Ram        string
	Hilos      int
	Ruta_juego string
}
