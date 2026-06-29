package config

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
		log.Fatalf("Unable to parse connection string: %s", err)
	}

	config.MaxConns = 25                       //maximo de connexiones abiertas
	config.MinConns = 1                        //Maintain a baseline of ready connections
	config.MaxConnLifetime = time.Hour         //Cierra conexiónes antiguas para memory leaks
	config.MaxConnIdleTime = 15 * time.Minute  //Cierra conneciones sitting idle
	config.HealthCheckPeriod = 1 * time.Minute //proactively ping the server

	db, err := pgxpool.NewWithConfig(ctx, config)
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

	q := `CREATE TABLE users3(
		user_id SERIAL PRIMARY KEY,
		name VARCHAR(255),
		card INT UNIQUE NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		is_active BOOLEAN NOT NULL DEFAULT TRUE,
		is_admin BOOLEAN NOT NULL DEFAULT FALSE,
		permissions VARCHAR(1) DEFAULT 'R'
	);`

	if _, err := db.Exec(ctx, q); err != nil {
		fmt.Println("error creato tabla users: ", err.Error())
	}
	return nil
}
