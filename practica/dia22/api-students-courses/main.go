package main

import (
	"api-students-courses/config"
	"api-students-courses/handlers"
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
		if err := config.MakeMigrations(); err != nil {
			log.Fatal(err)
		}
	}

	mux.HandleFunc("/students", handlers.StudentsHandlers)
	mux.HandleFunc("/courses", handlers.CoursesHandlers)
	mux.HandleFunc("/classes", handlers.StudentCoursesHandlers)
	mux.HandleFunc("/studentclasses", handlers.CourseStudentsHandlers)
	mux.HandleFunc("/teachers", handlers.TeacherHandlers)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	port := os.Getenv("PORT")
	http.ListenAndServe(port, mux)
}
