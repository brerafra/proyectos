/*
********************************************************

# Practia 3 - Dia 14

# Cliente http avanzado

petición get personalizada

Resultado: el html obtenidos usando un get requests
de la pagina example.com
*******************************************************
*/
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Payload struct {
	Mensaje string `json:"mensaje"`
}

func main() {
	//1. crear un cliente con un tiempo de espera
	cliente := &http.Client{
		Timeout: 10 * time.Second,
	}

	//2. Preparar los datos y la petición
	datos := Payload{Mensaje: "Hola desde Go"}
	jsonData, _ := json.Marshal(datos)

	req, err := http.NewRequest("POST", "https://example.com", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	//3. Añadir encabezados
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer my_token_secreto")

	//4. Ejecutar la peticion
	resp, err := cliente.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//5. leer la respuesta
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Código de estado: ", resp.Status)
	fmt.Println("Respuesta: ", string(body))
}
