package handlers

import (
	"encoding/json"
	"main/auth"
	"main/config"
	"main/domain"
	"main/pkg"
	"main/repository"
	"main/service"
	"net/http"
)

type TokenResponse struct {
	Token  string `json:"token"`
	Status string `json:"status"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	repo := repository.NewSQLUserRepository(config.GetConnection())
	service := service.NewUserService(repo)

	dbuser, err := service.GetUserByEmail(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	matchPassword := pkg.CheckPasswordHash(user.Password, dbuser.Password)
	if !matchPassword {
		http.Error(w, "Password incorrecto", http.StatusBadRequest)
		return
	}

	tokenString, err := auth.GenerateToken(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responde := TokenResponse{
		Token:  tokenString,
		Status: "ok",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responde)
}
