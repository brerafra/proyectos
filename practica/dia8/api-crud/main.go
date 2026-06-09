/*
*********************************************************************

# API dia 8

8.1 Api con Endpoints create/read/update/delete

Sección 1: se trabajan todos los endpoints y sus respuestas posibles
********************************************************************
*/
package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Status   bool
}

type Response struct {
	Status string `json:"status"`
}

var Usuarios []User

func deleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}

	queryParams := r.URL.Query()
	id := queryParams.Get("id")
	id_int, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "id invalido", http.StatusBadRequest)
		return
	}

	for i, u := range Usuarios {
		if u.Id == id_int {
			Usuarios[i].Status = false
			w.Header().Set("Content-Type", "application/json")
			response := Response{Status: "Usuario eliminado"}
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	http.Error(w, "usuario no existe", http.StatusBadRequest)

}

func updateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}
	var user User
	queryParams := r.URL.Query()
	id := queryParams.Get("id")
	id_int, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "id invalido", http.StatusBadRequest)
		return
	}
	//decodificar el cuerpo json del request
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, u := range Usuarios {
		if u.Id == id_int {
			if user.Name != "" {
				Usuarios[i].Name = user.Name
			}
			if user.Email != "" {
				Usuarios[i].Email = user.Email
			}
			if len(user.Password) > 8 {
				Usuarios[i].Password = user.Password
			}
			w.Header().Set("Content-Type", "application/json")
			response := Response{Status: "Usuario modificado"}
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	http.Error(w, "usuario no existe", http.StatusBadRequest)

}

func getUsers(w http.ResponseWriter, r *http.Request) {

	if err := json.NewEncoder(w).Encode(&Usuarios); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}

	var user User
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
		http.Error(w, "Password debe tener al menos 8 dígitos", http.StatusBadRequest)
		return
	}

	user.Status = true

	if len(Usuarios) == 0 {
		user.Id = 0
	} else {
		user.Id = Usuarios[len(Usuarios)-1].Id + 1
	}

	Usuarios = append(Usuarios, user)
	w.Header().Set("Content-Type", "application/json")
	response := Response{Status: "Usuario dado de alta"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", getUsers)
	mux.HandleFunc("POST /users", createUser)
	mux.HandleFunc("PUT /users", updateUser)
	mux.HandleFunc("DELETE /users", deleteUser)

	http.ListenAndServe(":8085", mux)
}
