package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secreto")

// Estructura para el cuerpo de la petición
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

// Estructura para la respuesta con los nuevos tokens
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func RefreshHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RefreshRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Body invalido", http.StatusBadRequest)
		return
	}

	//Analizar y validar el Refresh token
	token, err := jwt.Parse(req.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "refresh token invalido o expirado", http.StatusUnauthorized)
		return
	}

	//extraer los claims para asegurar que es un refresh token (opcional pero recomendado)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["type"] != "refresh" {
		http.Error(w, "token invalido", http.StatusUnauthorized)
		return
	}

	//extraer el subject (ejemp. ID usuario)
	userID := claims["sub"].(string)

	//Generar un nuevo Access token (corta duración: 15 min)
	accessTokenClaims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(15 * time.Minute).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	newAccessToken, _ := accessToken.SignedString(jwtKey)

	//Generar un nuevo refresh token (larga duración 7 dias)
	refreshTokenClaims := jwt.MapClaims{
		"sub":  userID,
		"type": "refresh",
		"exp":  time.Now().Add(7 * 24 * time.Hour).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	newRefreshToken, _ := refreshToken.SignedString(jwtKey)

	//Enviar los nuevos token al cliente

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(TokenResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	})
}

func main() {
	http.HandleFunc("/refresh", RefreshHandler)
	http.ListenAndServe(":8085", nil)
}
