package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

var jwtKey = []byte("tu_clave_secreta_super_segura")

func GenerarToken(userID uint, role string) (string, error) {
	//Expiration token
	expiratonTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiratonTime),
		},
	}

	//Crear el token con el algoritmo de firma
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//Firmar el token con la clave secreta
	return token.SignedString(jwtKey)

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

		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
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
