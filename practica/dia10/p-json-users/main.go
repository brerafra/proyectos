package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	var user User
	users, err := user.GetUsers()
	if err != nil {
		log.Fatal(err)
	}
	jsonData, _ := json.Marshal(users)
	os.WriteFile("users.json", jsonData, 00644)

	fmt.Println("Usuarios grabados en archivo.")
}
