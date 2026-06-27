package middleware

import (
	"main/auth"
	"net/http"
)

// JWTMiddleware restringe los enpoints a menos que el token sea valido
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Auhorization header missing", http.StatusUnauthorized)
			return
		}

		_, err := auth.ValidateToken(authHeader)
		if err != nil {
			http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
