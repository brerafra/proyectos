package handlers

import (
	"api-students-courses/config"
	"api-students-courses/domain"
	"api-students-courses/repository"
	"api-students-courses/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type PaginatedTeacherResponse struct {
	Data       []domain.Teacher `json:"data"`
	TotalRows  int              `json:"total_rows"`
	TotalPages int              `json:"total_pages"`
	Page       int              `json:"page"`
	Limit      int              `json:"limit"`
}

func CreateTeacherHandler(w http.ResponseWriter, r *http.Request) {
	var model domain.Teacher
	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	repo := repository.NewSQLTeacherRepository(config.GetConnection())
	service := service.NewTeacherService(repo)

	if err := service.RegisterTeacher(model.Name, model.Shift); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetTeacherByIdHandler(w http.ResponseWriter, r *http.Request) {
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

	repo := repository.NewSQLTeacherRepository(config.GetConnection())
	service := service.NewTeacherService(repo)

	teacher, err := service.GetTeacher(id)
	if err != nil {
		http.Error(w, "No se encontro usuario", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teacher)
}

func GetTeachersHandler(w http.ResponseWriter, r *http.Request) {
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

	repo := repository.NewSQLTeacherRepository(config.GetConnection())
	service := service.NewTeacherService(repo)

	teachers, totalRows, err := service.GetTeachers(page, limit)

	totalPages := (totalRows + limit - 1) / limit
	response := PaginatedTeacherResponse{
		Data:       teachers,
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Page:       page,
		Limit:      limit,
	}

	w.Header().Set("Content-Type", "applitation/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateTeacherHandler(w http.ResponseWriter, r *http.Request) {
	var teacher domain.Teacher

	if err := json.NewDecoder(r.Body).Decode(&teacher); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	repo := repository.NewSQLTeacherRepository(config.GetConnection())
	service := service.NewTeacherService(repo)

	if err := service.UpdateTeacher(teacher); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteTeacherHandler(w http.ResponseWriter, r *http.Request) {
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

	repo := repository.NewSQLTeacherRepository(config.GetConnection())
	service := service.NewTeacherService(repo)

	if err = service.DeleteTeacher(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func TeacherHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		idStr := r.URL.Query().Get("id")
		if idStr != "" {
			GetTeacherByIdHandler(w, r)
		} else {
			GetTeachersHandler(w, r)
		}
	case http.MethodPost:
		CreateTeacherHandler(w, r)
	case http.MethodPut:
		UpdateTeacherHandler(w, r)
	case http.MethodDelete:
		DeleteTeacherHandler(w, r)
	default:
		http.Error(w, "Metodo no permitido", http.StatusBadRequest)
		return

	}

}
