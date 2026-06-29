package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Extraer el encabezado Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Token de autorización faltante", http.StatusUnauthorized)
			return
		}

		//Separar el tipo de token (Beared y el token en si)
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Formato de token invalido", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]
		claims := &Claims{}

		//Analizar y validar el token

		err := godotenv.Load()
		if err != nil {
			log.Fatal("error loading .env file")
		}

		jwtKey := os.Getenv("JWT_KEY")

		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})

		if err != nil != !token.Valid {
			http.Error(w, "Token invalido o expirado", http.StatusUnauthorized)
			return
		}

		//Validar si el rol es 'admin'
		if claims.Role != "admin" {
			http.Error(w, "Acceso denegado: Se requieren privilegios de administrador", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
