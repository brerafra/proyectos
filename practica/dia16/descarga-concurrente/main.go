package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"https://httpbin.org",
		"https://httpbin.org",
		"https://httpbin.org",
	}

	//canal para recolectar los datos de forma segura
	resultados := make(chan string, len(urls))
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1) //Incrementamos el contador por cada gorutina

		go func(u string) {
			defer wg.Done() //Decrementa el contador al infalizar la gorutina

			resp, err := http.Get(u)
			if err != nil {
				resultados <- fmt.Sprintf("Error en %s, %v", u, err)
				return
			}
			defer resp.Body.Close()

			body, _ := io.ReadAll(resp.Body)
			resultados <- fmt.Sprintf("URL: %s | Tamaño: %d byes", u, len(body))
		}(url)
	}

	wg.Wait()
	close(resultados)

	//imprimimos los resultados
	for res := range resultados {
		fmt.Println(res)
	}
}
