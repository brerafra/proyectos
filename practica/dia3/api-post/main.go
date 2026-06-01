/*
*********************************************************************

practica día 3

1.1 Creación de api POST /users imprimiendo en terminal el post retornado

*********************************************************************
*/
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Users struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var user Users

	//decodificar el cuerpo json del request
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", createUser)
	http.ListenAndServe(":8085", mux)
}
