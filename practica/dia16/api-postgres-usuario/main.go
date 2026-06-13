package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func CreateTableQuery(p *pgxpool.Pool) {
	if _, err := p.Exec(context.Background(), "CREATE TABLE users (id SERIAL PRIMARY KEY, name VARCHAR(255) NOT NULL, email VARCHAR(255) UNIQUE NOT NULL);"); err != nil {
		log.Fatalf("Error while creating the table: %v\n", err)
	}
}

func InsertQuery() {
	db, ctx := GetConnection()
	if _, err := db.Exec(ctx, "insert into users(name, email) values($1,$2)", "John2", "johnysinsj@astronaut.com2"); err != nil {
		log.Fatalf("Error while inserting value into the table: %v\n", err)
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

	mux.HandleFunc("/users", usersHandlers)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	port := os.Getenv("PORT")
	http.ListenAndServe(port, mux)

}
