/*
********************************************************

# Practia 2 - Dia 14

# Cliente http basico

petición get básica

Resultado: el html obtenidos usando un get requests
de la pagina example.com
*******************************************************
*/
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://example.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
