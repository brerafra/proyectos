package pkg

import "golang.org/x/crypto/bcrypt"

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
