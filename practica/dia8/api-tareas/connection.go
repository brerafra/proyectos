/*
****************************************************************

connection.go -> en este paquete se obtiene el puntero a la estructura
de sql

packete hecho para sqlite3.

****************************************************************
*/
package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var db *sql.DB

func GetConnection() *sql.DB {
	if db != nil {
		return db
	}

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	db_name := os.Getenv("DB")

	db, err = sql.Open("sqlite3", db_name)
	if err != nil {
		panic(err)
	}

	return db
}

func MakeMigrations() error {
	db := GetConnection()

	q := `CREATE TABLE IF NOT EXISTS todos(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title VARCHAR(64) NULL,
			description VARCHAR(200) NULL,
			kind INTEGER NOT NULL,
			created_at TIMESTAMP DEFAULT DATETIME,
			updated_at TIMESTAMP NOT NULL
		);`
	_, err := db.Exec(q)
	if err != nil {
		return err
	}
	return nil

}
