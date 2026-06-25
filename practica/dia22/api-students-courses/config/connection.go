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
	q := `CREATE TABLE students(
		student_id SERIAL PRIMARY KEY,
		name VARCHAR(255)
	);`

	if _, err := db.Exec(ctx, q); err != nil {
		fmt.Println("error creando tabla students")
	}

	q = `CREATE TABLE courses(
		course_id SERIAL PRIMARY KEY,
		name VARCHAR(255)
	);`
	if _, err := db.Exec(ctx, q); err != nil {
		fmt.Println("error creando tabla courses")
	}

	q = `CREATE TABLE student_courses(
		student_course_id SERIAL PRIMARY KEY,
		student_id INT,
		course_id INT,
		CONSTRAINT fk_student_id FOREIGN KEY (student_id) REFERENCES students(student_id),
		CONSTRAINT fk_course_id FOREIGN KEY (course_id) REFERENCES courses(course_id),
		CONSTRAINT uc_student_id UNIQUE (student_id, course_id)
	);`

	if _, err := db.Exec(ctx, q); err != nil {
		fmt.Println("error creando tabla student_courses")
	}

	q = `CREATE TABLE teachers(
		studen_id SERIAL PRIMARY KEY,
		name VARCHAR(255),
		shift INT
	);`

	if _, err := db.Exec(ctx, q); err != nil {
		fmt.Println("error creando tabla teachers")
	}
	return nil
}
