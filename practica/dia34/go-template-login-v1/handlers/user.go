package handlers

import (
	"encoding/json"
	"main/config"
	"main/domain"
	"main/repository"
	"main/service"
	"net/http"
	"strconv"
)

type PaginatedUsersResponse struct {
	Data       []domain.User `json:"data"`
	TotalRows  int           `json:"total_rows"`
	TotalPages int           `json:"total_pages"`
	Page       int           `json:"page"`
	Limit      int           `json:"limit"`
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	repo := repository.NewSQLUserRepository(config.GetConnection())
	service := service.NewUserService(repo)

	if err := service.RegisterUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Query id es requerido", http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Query id debe ser un numero", http.StatusBadRequest)
		return
	}

	repo := repository.NewSQLUserRepository(config.GetConnection())
	service := service.NewUserService(repo)

	user, err := service.GetUser(id)
	if err != nil {
		http.Error(w, "No se encontro usuario", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	pageStr := query.Get("page")
	limitStr := query.Get("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	repo := repository.NewSQLUserRepository(config.GetConnection())
	service := service.NewUserService(repo)

	users, totalRows, err := service.GetUsers(page, limit)

	totalPages := (totalRows + limit - 1) / limit
	response := PaginatedUsersResponse{
		Data:       users,
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Page:       page,
		Limit:      limit,
	}

	w.Header().Set("Content-Type", "applitation/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user domain.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	repo := repository.NewSQLUserRepository(config.GetConnection())
	service := service.NewUserService(repo)

	if err := service.UpdateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Query id es requerido", http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Query id debe ser un numero", http.StatusBadRequest)
		return
	}
	repo := repository.NewSQLUserRepository(config.GetConnection())
	service := service.NewUserService(repo)

	if err = service.DeleteUser(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func UsersHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		idStr := r.URL.Query().Get("id")
		if idStr != "" {
			GetUserByIdHandler(w, r)
		} else {
			GetUsersHandler(w, r)
		}
	case http.MethodPost:
		CreateUserHandler(w, r)
	case http.MethodPut:
		UpdateUserHandler(w, r)
	case http.MethodDelete:
		DeleteUserHandler(w, r)
	default:
		http.Error(w, "Metodo no permitido", http.StatusBadRequest)
		return

	}

}
