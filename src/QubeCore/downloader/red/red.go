package red

import (
	"QbCore/versiones"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func Sha1File(path string) (string, error) {
	// obtiene el sha1 de un archivo del sistema
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func DownloadFile(url, destPath, expectedSHA1 string) error {

	if expectedSHA1 == "" && versiones.Existe_archivo(destPath) {

		return nil
	}

	// omite si ya existe
	if expectedSHA1 != "" { // esta seccion evita que se redescarguen archivos
		got, err := Sha1File(destPath)

		if err == nil && got == expectedSHA1 { //compara el archivo del sistema con el esperado, si es asi no hace nada
			return nil
		}
	}

	if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
		return err
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		return errors.New(fmt.Sprintf("la url no devolvio el codigo de estado esperado (200), codigo de estado retornado: %d", resp.StatusCode))

	}

	f, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return err
}
