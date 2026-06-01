/*
*********************************************************************

practica día 4

1.1 API POST /users donde se validará que exista un email y que el password enviado sea minimo de
8 dígitos

*********************************************************************
*/
package main

import (
	"encoding/json"
	"net/http"
)

type Users struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
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

	if user.Email == "" {
		http.Error(w, "Email requerido", http.StatusBadRequest)
		return
	}

	if len(user.Password) < 8 {
		http.Error(w, "Password debe ser igual o mayor de 8 caracteres", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", createUser)
	http.ListenAndServe(":8085", mux)
}
