package natives

import (
	"QbCore/versiones"
	"archive/zip"
	so "downloader/SO"
	"downloader/data"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// manejo de natives para versiones menores a la 1.19

// TODO: revisar todo esto

/*
pregunta si el archivo es necesario para el sistema operativo, si no es asi lo

ignora, emite un error que luego es ignorado
*/
func gestionar_natives(lib data.Library, SO string) (data.Artifact, error) {

	var error_ = errors.New("archivo no disponible")

	if !so.LibraryAllowed(lib) { // pregunta si esta habilitada
		return data.Artifact{}, error_
	}
	// buscar el classifier de windows
	classifier, ok := lib.Natives[SO] // clasifica por sistema op
	if !ok {
		return data.Artifact{}, error_
	}
	native, ok := lib.Downloads.Classifiers[classifier] //extrae el native

	if !ok || native.URL == "" {

		return data.Artifact{}, error_

	}

	return native, nil

}

func Maneja_Natives(tasks []data.Task, vj data.VersionJSON, OS string) []data.Task {

	// carga los natives y los agrega a la lista de tareas

	carpeta_natives := filepath.Join(versiones.Ruta_versiones, vj.ID, "natives")

	if err := os.MkdirAll(carpeta_natives, 0755); err != nil {
		fmt.Println("error al crear natives/: ", err)
	}

	for _, lib := range vj.Libraries {

		native, naterr := gestionar_natives(lib, OS)

		if naterr != nil {
			continue
		}

		tasks = append(tasks, data.Task{ //agregamos nueva task y retornamos
			URL:      native.URL,
			DestPath: filepath.Join(versiones.Ruta_minecraft, "libraries", filepath.FromSlash(native.Path)),
			SHA1:     native.SHA1,
			Label:    native.Path,
		})
	}
	return tasks
}

func Extraer_Natives(vj data.VersionJSON, OS string) error { // esta funcion extrae las natives, operacion necesaria para
	// versiones antiguas de Minecaft
	nativesDir := filepath.Join(versiones.Ruta_minecraft, "versions", vj.ID, "natives")
	os.MkdirAll(nativesDir, 0755)

	for _, lib := range vj.Libraries {

		native, naterr := gestionar_natives(lib, OS)

		if naterr != nil {
			continue
		}

		jarPath := filepath.Join(versiones.Ruta_minecraft, "libraries", filepath.FromSlash(native.Path))
		r, err := zip.OpenReader(jarPath)
		if err != nil {
			return fmt.Errorf("abriendo natives jar %s: %w", jarPath, err)
		}

		for _, f := range r.File {
			if strings.HasPrefix(f.Name, "META-INF") || f.FileInfo().IsDir() {
				continue
			}
			dest := filepath.Join(nativesDir, filepath.Base(f.Name))
			rc, _ := f.Open() //copia los distintos archivos
			out, _ := os.Create(dest)
			_, err := io.Copy(out, rc)
			if err != nil {

				return fmt.Errorf("copiando %s, %w", f.Name, err)

			}

			out.Close()
			rc.Close()
		}
		defer r.Close()
	}
	return nil
}
