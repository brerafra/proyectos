package main

import (
	"flag"
	"log"
	"log/slog"
	"main/config"
	"main/handlers"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	slog.SetLogLoggerLevel(slog.LevelDebug)
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

	mux.HandleFunc("/users", handlers.UsersHandlers)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	port := os.Getenv("PORT")

	slog.Info("Iniciando proceso en puerto" + port)

	http.ListenAndServe(port, mux)
}
