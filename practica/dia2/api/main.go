/*
*********************************************************************

practica día 2

1.1 Creación de api GET /users con respuesta en json con todos los usuarios (en codigo duro)

	se puede testear usando curl http://localhost:8085/users

*********************************************************************
*/
package main

import (
	"encoding/json"
	"net/http"
)

type Users struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	Usuarios := []Users{
		{1, "Brenthon"},
		{2, "Manuel"},
		{3, "Israel"},
		{4, "Irma"},
		{5, "Zohar"},
		{6, "Iris"},
		{7, "Jo M"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Usuarios)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", getUsers)
	http.ListenAndServe(":8085", mux)
}
