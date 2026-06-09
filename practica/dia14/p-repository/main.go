package main

import (
	"database/sql"
	"log"
	"main/repository"
	"main/service"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	q := `CREATE TABLE IF NOT EXISTS users(id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, email TEXT)`
	_, err = db.Exec(q)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewSQLRepository(db)
	UserService := service.NewUserService(userRepo)

	newUser, err := UserService.RegisterUSer("Juan Pérez", "juan@example.com")
	if err != nil {
		log.Println("Error al crear usuario: ", err)
	} else {
		log.Printf("USuario creado con ID: %d\n", newUser.ID)
	}

}
