package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo")
}

func GetNotesHandler(w http.ResponseWriter, r *http.Request) {
	n := new(Note)

	notes, err := n.GetNotes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	j, err := json.Marshal(notes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func CreateNotesHandler(w http.ResponseWriter, r *http.Request) {
	var note Note

	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = note.create()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UpdateNotesHandler(w http.ResponseWriter, r *http.Request) {
	var note Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = note.Update()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteNotesHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Query id es requerido", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Query id debe de ser un número", http.StatusBadRequest)
		return
	}

	var note Note

	err = note.delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func NotesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetNotesHandler(w, r)
	case http.MethodPost:
		CreateNotesHandler(w, r)
	case http.MethodPut:
		UpdateNotesHandler(w, r)
	case http.MethodDelete:
		DeleteNotesHandler(w, r)
	default:
		http.Error(w, "Metodo no permitido", http.StatusBadRequest)
		return
	}
}

func main() {
	mux := http.NewServeMux()

	migrate := flag.Bool(
		"migrate", false, "Crea las tablas en la base de datos",
	)
	flag.Parse()
	if *migrate {
		if err := MakeMigrations(); err != nil {
			log.Fatal(err)
		}
	}
	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/notes", NotesHandler)
	http.ListenAndServe(":8085", mux)
}
