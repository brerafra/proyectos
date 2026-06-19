package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"main.go/config"
	"main.go/domain"
	"main.go/repository"
	"main.go/service"
)

type PaginatedEmployeeResponse struct {
	Data       []domain.Empleado `json:"data"`
	TotalRows  int               `json:"total_rows"`
	TotalPages int               `json:"total_pages"`
	Page       int               `json:"page"`
	Limit      int               `json:"limit"`
}

func CreateEmpHandler(w http.ResponseWriter, r *http.Request) {
	var e domain.Empleado

	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	repo := repository.NewSQLEmployeeRepository(config.GetConnection())
	service := service.NewEmployeeService(repo)

	if err := service.RegisterEmployee(e); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetEmployeeByIdHandler(w http.ResponseWriter, r *http.Request) {
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

	repo := repository.NewSQLEmployeeRepository(config.GetConnection())
	service := service.NewEmployeeService(repo)

	jdata, err := service.GetById(id)
	if err != nil {
		http.Error(w, "No se encontro usuario", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jdata)
}

func GetEmployeesHandler(w http.ResponseWriter, r *http.Request) {
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

	repo := repository.NewSQLEmployeeRepository(config.GetConnection())
	service := service.NewEmployeeService(repo)

	empleados, totalRows, err := service.GetAll(page, limit)
	if err != nil {
		http.Error(w, "No se encontro empleados", http.StatusBadRequest)
		return
	}

	totalP := (totalRows + limit - 1) / limit
	response := PaginatedEmployeeResponse{
		Data:       empleados,
		TotalRows:  totalRows,
		TotalPages: totalP,
		Page:       page,
		Limit:      limit,
	}

	w.Header().Set("Content-Type", "applitation/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	var e domain.Empleado

	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	repo := repository.NewSQLEmployeeRepository(config.GetConnection())
	service := service.NewEmployeeService(repo)

	if err := service.Update(e); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteEmployeeHandler(w http.ResponseWriter, r *http.Request) {
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

	repo := repository.NewSQLEmployeeRepository(config.GetConnection())
	service := service.NewEmployeeService(repo)

	if err = service.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func EmployeeHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		idStr := r.URL.Query().Get("id")
		if idStr != "" {
			GetEmployeeByIdHandler(w, r)
		} else {
			GetEmployeesHandler(w, r)
		}
	case http.MethodPost:
		CreateEmpHandler(w, r)
	case http.MethodPut:
		UpdateEmployeeHandler(w, r)
	case http.MethodDelete:
		DeleteEmployeeHandler(w, r)
	default:
		http.Error(w, "Metodo no permitido", http.StatusBadRequest)
		return

	}
}
