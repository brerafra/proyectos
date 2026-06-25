package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	//Convert password string to byte slice and apply default cost(10)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares a plaintext password with its hashed version
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	password := "mySecretPassword123"

	//1. Hash the password
	hash, err := HashPassword(password)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	fmt.Println("Password:", password)
	fmt.Println("Hash:	", hash)

	//2. Verify a correct password
	match := CheckPasswordHash(password, hash)
	fmt.Println("Macht (correct):", match)

	//3. Verify an incorrect password
	wrongMatch := CheckPasswordHash("wrongPassword", hash)
	fmt.Println("Match (wrong);", wrongMatch)
}
