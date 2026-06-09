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
	db, err := sql.Open("sqlite3", db_name)
	if err != nil {
		panic(err)
	}

	return db
}

func MakeMigrations() error {
	db := GetConnection()
	q := `CREATE TABLE IF NOT EXISTS users(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(200) NOT NULL,
			email VARCHAR(200) UNIQUE,
			pin VARCHAR(4) NOT NULL,
			status INTEGER NOT NULL
		);`
	_, err := db.Exec(q)
	if err != nil {
		return err
	}
	return nil
}
