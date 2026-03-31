package goruntinas

import (
	"QbCore/consola"
	"downloader/data"
	"downloader/red"
	"fmt"
	"sync"

	"github.com/pterm/pterm"
)

// este modulo maneja la concurrencia del downloader

func worker(wg *sync.WaitGroup, errores chan string, ch chan data.Task, carga *pterm.ProgressbarPrinter) {

	defer wg.Done()
	for task := range ch {
		err := red.DownloadFile(task.URL, task.DestPath, task.SHA1)

		if err != nil {
			errores <- fmt.Sprintf("FALLO [%s]: %v", task.Label, err) // recolecta errores en caso de haberlo
			continue
		}

		carga.Increment()
	}

}

// funcion auxiliar para runworkers
func lanzar_goruntina(wg *sync.WaitGroup, errores chan string, workers int, ch chan data.Task, carga *pterm.ProgressbarPrinter) {

	for i := 0; i < workers; i++ {
		wg.Add(1)

		go worker(wg, errores, ch, carga)
	}

	wg.Wait()
	carga.Stop()
	close(errores) // espera hasta la ultima goruntina y cierra el canal, en ese orde NO OLVIDAR
	// quien escribe el canal? las goruntinas , cuando terminan? primero debo esperar (wg.wait) y luego cierro el canal
	fmt.Print("\n")

	if len(errores) > 0 { //imprime los errores que recolecto
		fmt.Printf("\n%d error/es:\n", len(errores))
		for e := range errores {
			fmt.Println(" ", e)
		}
	}

}

// descarga los archivos con concurrencia para agilizar
func RunWorkers(tasks []data.Task, workers int) {
	var total_tareas = len(tasks)

	ch := make(chan data.Task, total_tareas)
	errores := make(chan string, total_tareas)
	for _, t := range tasks { // pasa las tareas
		ch <- t
	}
	close(ch)

	var wg = sync.WaitGroup{}

	carga := consola.Crear_barra(total_tareas, "(+) Preparando lanzamiento... ")
	lanzar_goruntina(&wg, errores, workers, ch, carga)
}
