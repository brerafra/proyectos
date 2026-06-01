/*
*********************************************************************

# Proyecto día 6

5.2 Practica para leer y escribir un archivo js.
*****************************************************************
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Persona struct {
	Nombre string `json:"nombre"`
	Edad   int    `json:"edad"`
}

func main() {
	p := []Persona{{"Ana", 30}, {"Irma", 33}, {"Leo", 32}, {"Faby", 31}}
	//Convertir struct a json
	jsonData, _ := json.Marshal(p)

	os.WriteFile("datos.js", jsonData, 0644)

	data, err := ioutil.ReadFile("datos.js")
	if err != nil {
		log.Fatal(err)
	}

	var people []Persona

	err = json.Unmarshal(data, &people)
	if err != nil {
		log.Fatal(err)
	}

	for _, k := range people {
		fmt.Printf("Usuario: %s, Edad: %d\n", k.Nombre, k.Edad)
	}

	//

}
