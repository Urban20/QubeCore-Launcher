package goruntinas

import (
	"downloader/data"
	"downloader/red"
	"fmt"
	"sync"
	"sync/atomic"
)

// este modulo maneja la concurrencia del downloader

// funcion auxiliar para runworkers
func lanzar_goruntina(wg *sync.WaitGroup, done *atomic.Int64, errors []string, workers int, total int64, ch chan data.Task, mu *sync.Mutex) {
	for i := 0; i < workers; i++ {
		wg.Add(1)

		go func() {

			defer wg.Done()
			for task := range ch {
				err := red.DownloadFile(task.URL, task.DestPath, task.SHA1)

				n := done.Add(1)
				mu.Lock()

				if err != nil {
					errors = append(errors, fmt.Sprintf("FALLO [%s]: %v", task.Label, err))
					fmt.Printf("\r[%d/%d] ✗ %s\n", n, total, task.Label)
					continue
				}

				fmt.Printf("\r[%d/%d] ✓ %-60s", n, total, task.Label)

				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println()

	if len(errors) > 0 {
		fmt.Printf("\n%d error(s):\n", len(errors))
		for _, e := range errors {
			fmt.Println(" ", e)
		}
	}

}

// descarga los archivos con concurrencia para agilizar
func RunWorkers(tasks []data.Task, workers int) {
	ch := make(chan data.Task, len(tasks))
	for _, t := range tasks {
		ch <- t
	}
	close(ch)

	var (
		wg     sync.WaitGroup
		done   atomic.Int64
		total  = int64(len(tasks))
		mu     sync.Mutex
		errors []string
	)
	lanzar_goruntina(&wg, &done, errors, workers, total, ch, &mu)
}
