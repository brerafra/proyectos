package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// Lanzamos exactamente 10 goroutines
	for i := 1; i <= 10; i++ {
		wg.Add(1) //Añadimos una tarea al contador

		go func(id int) {
			defer wg.Done() // Restamos una tarea al finalizar la goroutina
			fmt.Println("Goroutine: ", id, " ejecutandose...")
		}(i)
	}
	wg.Wait() //esperamos a que el contador de WaitGroup llegue a 0
	fmt.Println("Todas las goroutines han finalizado.")
}
