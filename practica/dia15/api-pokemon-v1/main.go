package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Pokemon representa la estructura de los datos que vamos a consumir de la API.
type Pokemon struct {
	Name   string `json:"name"`
	Weight int    `json:"weight"`
}

func main() {
	//1. Definir la URL del endpoint de la API
	url := "https://pokeapi.co/api/v2/pokemon/pikachu"

	//2.- Crear el cliente http con un tiempo de espera (timeout)
	client := &http.Client{Timeout: 10 * time.Second}

	//3.- Crear la petición GET
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error al crear la petición: ", err)
		os.Exit(1)
	}

	//4.- Ejecutar la peticion
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error al realizar la petición: %v", err)
	}

	defer resp.Body.Close()

	//5. Validar la respuesta sea exitosa (codigo 200)
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: estado de la respuesta %v", resp.Status)
	}

	//6. Decodificar el json de la respuesta en nuestra struct
	var pokemon Pokemon
	err = json.NewDecoder(resp.Body).Decode(&pokemon)
	if err != nil {
		log.Fatalf("Error: al decodificar %v", err)
	}

	fmt.Printf("Pokemon encontrado!\nNombre: %s\nPeso: %d hectogramos\n", pokemon.Name, pokemon.Weight)
}
