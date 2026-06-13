package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func getUrl(url string) ([]byte, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, fmt.Errorf("Error al crear la petición: %v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, fmt.Errorf("Error al realizar la petición: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("Error: estado de la respuesta %v", resp.Status)
	}

	cuerpoBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Error leyendo el cuerpo %v", err)
	}

	return cuerpoBytes, nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	url_api := os.Getenv("URL_API")

	body, err := getUrl(url_api)
	if err != nil {
		log.Fatalf("Error en get: %v", err.Error())
	}

	var pokemonGet PokemonGet
	err = json.Unmarshal(body, &pokemonGet)
	if err != nil {
		log.Fatalf("Error: al decodificar %v", err)
	}

	fmt.Printf("Pokemons encontrados!\nCantidad:%d\nSiguiente url:%s\nAnterior url:%s\n", pokemonGet.Count, pokemonGet.Next, pokemonGet.Previous)

	for _, k := range pokemonGet.Results {
		//fmt.Printf("%d, Pokemon: %s, url:%s\n", i, k.Name, k.Url)
		body, err = getUrl(k.Url)
		if err != nil {
			log.Fatalf("Error en get: %v", err.Error())
		}

		var p Pokemon
		err = json.Unmarshal(body, &p)
		if err != nil {
			log.Fatalf("Error: al decodificar %v", err)
		}
		fmt.Printf("%d,Nombre: %s, Peso:%d\n", p.Id, p.Name, p.Weight)
		for j, a := range p.Abilities {
			fmt.Printf("hability: %d, name:%s\n", j, a.Ability.Name)
		}
	}
}
