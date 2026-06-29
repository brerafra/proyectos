/***
	Validar disponibilidad de rutas usando worker pools Golang


***/

package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Estructurtura para el trabajo (input)
type Job struct {
	ID    int
	Route string
}

// estructura para el resultado (output)
type Result struct {
	Job        Job
	StatusCode int
	StatusText string
}

// Funcion worker que procesa cada ruta
func worker(id int, jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		fmt.Printf("[Worker %d] Iniciando validación para: %s\n", id, job.Route)

		//simulamos la petición http
		resp, err := http.Get(job.Route)
		if err != nil {
			results <- Result{
				Job:        job,
				StatusCode: 0,
				StatusText: fmt.Sprintf("Error: %s", err.Error()),
			}
			continue
		}
		//Cerramos el cuerpo de la respuesta para evitar fugas de memoria
		resp.Body.Close()

		//Enviamos el resultado al canal
		results <- Result{
			Job:        job,
			StatusCode: resp.StatusCode,
			StatusText: resp.Status,
		}
	}
}

func main() {
	routes := []string{
		"https://google.com",
		"https://github.com",
		"https://httpbin.org",             //ruta simulada no encontrada
		"https://httpbin.org",             //ruta simulada lenta
		"https://invalid-url-example.com", //ruta errónea
	}

	numJobs := len(routes)
	numWorkers := 3

	//Canales para gestionar el flujo
	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)
	var wg sync.WaitGroup

	//1. Lanzamos los workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(workerId int) {
			defer wg.Done()
			worker(workerId, jobs, results)
		}(w)
	}

	//2. Enviamos las rutas al canal de trabajos (jobs)
	for i, route := range routes {
		jobs <- Job{ID: i + 1, Route: route}
	}
	close(jobs) //Cerramos para indicar que ya nos e enviarán mas trabajos

	//3. Esperamos a que todos los workers terminen en una goroutine separada
	go func() {
		wg.Wait()
		close(results) //Cerramos resultados una vez procesados todos
	}()

	//4. Imprimimos los resultados a medida que llegan
	for res := range results {
		if res.StatusCode == 200 {
			fmt.Printf(" [OK] %s -> %s\n", res.Job.Route, res.StatusText)
		} else {
			fmt.Printf(" [FAIL] %s -> Código: %d (%s)\n", res.Job.Route, res.StatusCode, res.StatusText)
		}
	}
}
