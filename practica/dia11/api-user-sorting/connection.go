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
	if err = godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	db_name := os.Getenv("DB")
	db, err := sql.Open("sqlite3", db_name)
	if err != nil {
		panic(err)
	}
	return db
}

func MakeMigrations() error {
	db := GetConnection()
	q := `CREATE TABLE IF NOT EXISTS productos(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			nombre VARCHAR(64) NOT NULL,
			precio INTEGER NOT NULL,
			status BOOLEAN NOT NULL CHECK(status IN (0, 1))
		);`

	_, err := db.Exec(q)
	if err != nil {
		return err
	}
	return nil
}
