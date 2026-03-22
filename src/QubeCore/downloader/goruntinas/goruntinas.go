package goruntinas

import (
	"downloader/data"
	"downloader/red"
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/pterm/pterm"
)

// este modulo maneja la concurrencia del downloader

func worker(wg *sync.WaitGroup, errores chan string, ch chan data.Task, conteo *atomic.Int64, total_carga int) {

	defer wg.Done()
	for task := range ch {
		err := red.DownloadFile(task.URL, task.DestPath, task.SHA1)

		//n := done.Add(1) // hace un conteo sincronizado, si lo quiero hacer con i++ simple no sigue sucesion
		conteo.Add(1)
		pterm.BgGreen.Printf("\r(+) [%d/%d]", total_carga, conteo.Load())

		if err != nil {
			errores <- fmt.Sprintf("FALLO [%s]: %v", task.Label, err) // recolecta errores en caso de haberlo

			continue
		}

	}

}

// funcion auxiliar para runworkers
func lanzar_goruntina(wg *sync.WaitGroup, errores chan string, workers int, ch chan data.Task, conteo *atomic.Int64, total_tareas int) {

	for i := 0; i < workers; i++ {
		wg.Add(1)

		go worker(wg, errores, ch, conteo, total_tareas)
	}

	wg.Wait()
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
	var conteo atomic.Int64

	lanzar_goruntina(&wg, errores, workers, ch, &conteo, total_tareas)
}
