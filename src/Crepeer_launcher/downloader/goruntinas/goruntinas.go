package goruntinas

import (
	"downloader/data"
	"downloader/red"
	"fmt"
	"sync"
	"sync/atomic"
)

// este modulo maneja la concurrencia del downloader

func worker(wg *sync.WaitGroup, errores chan string, done *atomic.Int64, total int64, ch chan data.Task) {

	defer wg.Done()
	for task := range ch {
		err := red.DownloadFile(task.URL, task.DestPath, task.SHA1)

		n := done.Add(1) // hace un conteo sincronizado, si lo quiero hacer con i++ simple no sigue sucesion

		if err != nil {
			errores <- fmt.Sprintf("FALLO [%s]: %v", task.Label, err) // recolecta errores en caso de haberlo
			fmt.Printf("\r[%d/%d] ✗ %s\n", n, total, task.Label)
			continue
		}

		fmt.Printf("\r[%d/%d] ✓ %-60s", n, total, task.Label)

	}

}

// funcion auxiliar para runworkers
func lanzar_goruntina(wg *sync.WaitGroup, done *atomic.Int64, errores chan string, workers int, total int64, ch chan data.Task) {

	for i := 0; i < workers; i++ {
		wg.Add(1)

		go worker(wg, errores, done, total, ch)
	}

	wg.Wait()
	close(errores) // espera hasta la ultima goruntina y cierra el canal, en ese orde NO OLVIDAR
	// quien escribe el canal? las goruntinas , cuando terminan? primero debo esperar (wg.wait) y luego cierro el canal

	if len(errores) > 0 { //imprime los errores que recolecto
		fmt.Printf("\n%d error(s):\n", len(errores))
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

	var (
		wg    sync.WaitGroup
		done  atomic.Int64        //lleva un conteo de las tareas realizadas
		total = int64(len(tasks)) //cuenta las tareas totales
	)
	lanzar_goruntina(&wg, &done, errores, workers, total, ch)
}
