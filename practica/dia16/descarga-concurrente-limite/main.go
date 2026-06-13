package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"https://example.com",
		"https://example.com",
		"https://example.com",
		"https://example.com",
		"https://example.com",
		"https://example.com",
		"https://example.com",
		"https://example.com",
	}

	//limite de 2 descargas concurrentes simultáneas
	limite := make(chan struct{}, 2)
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)

		//enviar señal al canal; esto bloque si el canal ya tiene 2 descargas
		limite <- struct{}{}

		go func(u string) {
			defer wg.Done()
			defer func() {
				<-limite
			}()

			fmt.Printf("Iniciando la desarga de :%s\n", u)
			resp, err := http.Get(u)
			if err != nil {
				fmt.Printf("Error al descargar %s, %v\n", u, err)
				return
			}
			defer resp.Body.Close()
			fmt.Printf("Descarga completa:%s(Código: %d)\n", u, resp.StatusCode)
		}(url)
	}
	wg.Wait()
}
