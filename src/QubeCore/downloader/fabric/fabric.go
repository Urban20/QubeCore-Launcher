package fabric

import (
	"QbCore/versiones"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var Loader = filepath.Join(
	versiones.Ruta_minecraft,
	"libraries",
	"net", "fabricmc", "fabric-loader",
	"0.18.5",
	"fabric-loader-0.18.5.jar",
)

type Library_fabric struct {
	Sha1   string
	Sha256 string
	Size   int
	Name   string
	Url    string
	Md5    string
}

type Fabric struct {
	InheritsFrom string

	ReleaseTime string

	Mainclass string

	Libraries []Library_fabric
}

func Formatear_json_fabric() Fabric {

	// esto esta hardcodeado, si funciona hay que cambiarlo
	fabric := Fabric{}
	arch, _ := os.Open(filepath.Join(versiones.Ruta_versiones, "fabric-loader-0.18.5-1.21.11", "fabric-loader-0.18.5-1.21.11.json"))
	b, _ := io.ReadAll(arch)

	json.Unmarshal(b, &fabric)

	return fabric

}

func Obtener_librerias_fabric(fabric Fabric) []string {

	// lo que hice aca es formatear con rutamaven y
	// luego añadir una iteracion para guardar las rutas en una lista

	librerias := []string{}

	for _, libreria := range fabric.Libraries {

		nombre := RutaMaven(libreria.Name)

		librerias = append(librerias, nombre)

	}

	return librerias //tiene una lista con las rutas

}

// formatea las rutas de las librerias de fabric en disco
func RutaMaven(archivo string) string {
	partes := strings.Split(archivo, ":")
	if len(partes) < 3 {
		return ""
	}
	grupo := strings.ReplaceAll(partes[0], ".", "/")
	artifact := partes[1]
	version := partes[2]
	//return group + "/" + artifact + "/" + version + "/" + artifact + "-" + version + ".jar"
	return filepath.Join(versiones.Ruta_minecraft, "libraries", grupo, artifact, version, fmt.Sprintf("%s-%s.jar", artifact, version))

}

func Iniciar_sistema_fabric() []string {
	fabric := Formatear_json_fabric()

	return Obtener_librerias_fabric(fabric)

}
