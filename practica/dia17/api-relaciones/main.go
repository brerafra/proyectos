package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
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

	mux.HandleFunc("/users", usersHandlers)
	mux.HandleFunc("/tasks", tasksHandlers)
	mux.HandleFunc("/utasks", userTaskHandlers)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	port := os.Getenv("PORT")
	http.ListenAndServe(port, mux)

}
