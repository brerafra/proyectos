/*
*********************************************************************

practica día 1

1.1 Creación de api GET /health con respuesta en json:

	{
		"status":"ok"
	}

	se puede testear usando curl http://localhost:8085/health

*********************************************************************
*/
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ResponseHealth struct {
	Status string `json:"status"`
}

type ResponseDateTime struct {
	Tiempo string `json:"tiempo"`
}

func getHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	status := ResponseHealth{}
	status.Status = "ok"
	json.NewEncoder(w).Encode(status)
}

func getDateTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	t := time.Now()
	fmt.Println("Tiempo completo: ", t)

	formato := t.Format("2006-01-02 15:04:05")
	fmt.Println("Formato personalizado: ", formato)

	anio, mes, dia := t.Date()
	hora, min, seg := t.Clock()
	fmt.Printf("Fecha: %d/%d/%d\n", anio, mes, dia)
	fmt.Printf("Hora: %02d:%02d:%02d\n", hora, min, seg)

	DateTime := ResponseDateTime{Tiempo: formato}
	json.NewEncoder(w).Encode(DateTime)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", getHealth)
	mux.HandleFunc("GET /date", getDateTime)

	http.ListenAndServe(":8085", mux)
}
