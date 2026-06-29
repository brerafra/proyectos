package main

import "fmt"

func main() {
	//1. Crear un canal que transporta datos de tipo string
	messages := make(chan string)

	//2. Enviar datos al canal desde una Goroutine anónima
	go func() {
		messages <- "¡Hola desde el canal!"
	}()

	//3. Recibir datos del canal desde la Goroutine principal
	msg := <-messages //operación recepción
	fmt.Println(msg)
}
