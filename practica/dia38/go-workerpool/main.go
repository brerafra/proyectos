package main

import (
	"fmt"
	"time"
)

// Worker procesa los trabajos de la cola
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d inició el trabajo %d \n", id, job)

		//simulamos un tiempo de procesamiento (ej. 1 segundo)
		time.Sleep(1 * time.Second)

		fmt.Printf("Worker %d finalizó el trabajo %d\n", id, job)
		results <- job * 2 //Enviamos el resultado al canal
	}
}

func main() {
	const numJobs = 5
	const numWorkers = 3

	//Canales para trabajos y resultados
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	//1. Iniciamos el gropo de workes (goroutines)
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	//2. Enviamos los trabajos a la cola y cerramos el canal
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) //le decimos a los workers que ya no hay mas trabajos

	//3. Recogemos todos los resultados
	for a := 1; a <= numJobs; a++ {
		<-results //esperamos que cada worker termine
	}

	fmt.Println("todos los trabajos fueron procesados exitosamente.")
}
