package goruntinas

import (
	"downloader/data"
	"downloader/red"
	"fmt"
	"sync"
)

// este modulo maneja la concurrencia del downloader

func worker(wg *sync.WaitGroup, errores chan string, ch chan data.Task) {

	defer wg.Done()
	for task := range ch {
		err := red.DownloadFile(task.URL, task.DestPath, task.SHA1)

		//n := done.Add(1) // hace un conteo sincronizado, si lo quiero hacer con i++ simple no sigue sucesion

		if err != nil {
			errores <- fmt.Sprintf("FALLO [%s]: %v", task.Label, err) // recolecta errores en caso de haberlo

			continue
		}

	}

}

// funcion auxiliar para runworkers
func lanzar_goruntina(wg *sync.WaitGroup, errores chan string, workers int, ch chan data.Task) {

	for i := 0; i < workers; i++ {
		wg.Add(1)

		go worker(wg, errores, ch)
	}

	wg.Wait()
	close(errores) // espera hasta la ultima goruntina y cierra el canal, en ese orde NO OLVIDAR
	// quien escribe el canal? las goruntinas , cuando terminan? primero debo esperar (wg.wait) y luego cierro el canal

	if len(errores) > 0 { //imprime los errores que recolecto
		fmt.Printf("\n%d error/es:\n", len(errores))
		for e := range errores {
			fmt.Println(" ", e)
		}
	}

}

// descarga los archivos con concurrencia para agilizar
func RunWorkers(tasks []data.Task, workers int) {
	ch := make(chan data.Task, len(tasks))
	errores := make(chan string, len(tasks))
	for _, t := range tasks { // pasa las tareas
		ch <- t
	}
	close(ch)

	var wg = sync.WaitGroup{}

	lanzar_goruntina(&wg, errores, workers, ch)
}
