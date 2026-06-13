package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var db *pgxpool.Pool

func GetConnection() (*pgxpool.Pool, context.Context) {
	ctx := context.Background()
	if db != nil {
		return db, ctx
	}

	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	url_db := os.Getenv("URL_DB")
	config, err := pgxpool.ParseConfig(url_db)
	if err != nil {
		log.Fatalf("Unable to parse connection string: %v", err)
	}

	config.MaxConns = 25                       //maximo de connexiones abiertas
	config.MinConns = 1                        //Maintain a baseline of ready connections
	config.MaxConnLifetime = time.Hour         //Cierra conexiónes antiguas para memory leaks
	config.MaxConnIdleTime = 15 * time.Minute  //Cierra conneciones sitting idle
	config.HealthCheckPeriod = 1 * time.Minute //proactively ping the server

	db, err = pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatal("Error While creatiing connection to the database!!")
	}

	connection, err := db.Acquire(ctx)
	if err != nil {
		log.Fatal("Error while acquring connection from the database pool!")
	}

	if err = connection.Ping(ctx); err != nil {
		log.Fatal("Coud not ping database")
	}

	return db, ctx
}

func MakeMigrations() error {
	db, ctx := GetConnection()
	q := `CREATE TABLE users(
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			is_active BOOLEAN NOT NULL DEFAULT TRUE,
			is_admin BOOLEAN NOT NULL DEFAULT FALSE
		);`

	if _, err := db.Exec(ctx, q); err != nil {
		//return err
	}

	q = `CREATE TABLE posts(
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL,
			title VARCHAR(100) NULL,
			content VARCHAR(255) NULL
		);`

	if _, err := db.Exec(ctx, q); err != nil {
		fmt.Println("error en post")
		//return err
	}

	q = `CREATE TABLE tasks(
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL,
			title VARCHAR(100),
			content VARCHAR(255),
			kind INTEGER NOT NULL,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP NULL
		);`

	if _, err := db.Exec(ctx, q); err != nil {
		fmt.Println("error en tasks")
		//return err
	}

	return nil
}
