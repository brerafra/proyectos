package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"main.go/config"
	"main.go/handlers"
)

func main() {
	mux := http.NewServeMux()
	migrate := flag.Bool(
		"migrate", false, "Crea las tablas en la base de datos",
	)
	flag.Parse()
	if *migrate {
		if err := config.MakeMigrations(); err != nil {
			log.Fatal(err)
		}
	}

	mux.HandleFunc("/employees", handlers.EmployeeHandlers)

	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	port := os.Getenv("PORT")
	http.ListenAndServe(port, mux)

}
