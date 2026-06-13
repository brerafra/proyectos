/*
********************************************************

# Practia 4 - Dia 14

# Cliente http basico

petición get básica con context

Resultado: el html obtenidos usando un get requests
de la pagina example.com
*******************************************************
*/
package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://api.example.com/data", nil)
	if err != nil {
		log.Fatalf("Request Creation failed: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	log.Printf("Status: %s", resp.Status)
}
