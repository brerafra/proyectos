package auth

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type CustomClaims struct {
	Username    string `json:"username"`
	Role        string `json:"role"`
	Permissions string `json:"permissions"`
	jwt.RegisteredClaims
}

func GenerateToken(username string, role string, permissions string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)

	claims := &CustomClaims{
		Username:    username,
		Role:        role,
		Permissions: permissions,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	jwtKey := os.Getenv("JWT_KEY")
	fmt.Println(jwtKey)
	return token.SignedString([]byte(jwtKey))
}

func ValidateToken(tokenStr string) (*CustomClaims, error) {
	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		err := godotenv.Load()
		if err != nil {
			log.Fatal("error loading .env file")
		}

		jwtKey := os.Getenv("JWT_KEY")
		return []byte(jwtKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token invalido")
	}

	return claims, nil
}
