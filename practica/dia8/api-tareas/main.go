package main

import (
	"flag"
	"log"
	"net/http"
)

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
	mux.HandleFunc("/todos", TodosHandlers)
	http.ListenAndServe(":8085", mux)
}
