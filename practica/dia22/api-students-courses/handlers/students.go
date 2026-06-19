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

type PaginatedStudentResponse struct {
	Data       []domain.Student `json:"data"`
	TotalRows  int              `json:"total_rows"`
	TotalPages int              `json:"total_pages"`
	Page       int              `json:"page"`
	Limit      int              `json:"limit"`
}

func CreateStudentHandler(w http.ResponseWriter, r *http.Request) {
	var student domain.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	studentRepo := repository.NewSQLRepository(config.GetConnection())
	studentService := service.NewStudentService(studentRepo)

	if _, err := studentService.RegisterStudent(student.Name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetStudentByIdHandler(w http.ResponseWriter, r *http.Request) {
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

	studentRepo := repository.NewSQLRepository(config.GetConnection())
	studentService := service.NewStudentService(studentRepo)

	student, err := studentService.GetStudent(id)
	if err != nil {
		http.Error(w, "No se encontro usuario", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func GetStudentsHandler(w http.ResponseWriter, r *http.Request) {
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

	studentRepo := repository.NewSQLRepository(config.GetConnection())
	studentService := service.NewStudentService(studentRepo)

	students, totalRows, err := studentService.GetStudents(page, limit)

	totalPages := (totalRows + limit - 1) / limit
	response := PaginatedStudentResponse{
		Data:       students,
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Page:       page,
		Limit:      limit,
	}

	w.Header().Set("Content-Type", "applitation/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateStudentHandler(w http.ResponseWriter, r *http.Request) {
	var student domain.Student

	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	studentRepo := repository.NewSQLRepository(config.GetConnection())
	studentService := service.NewStudentService(studentRepo)

	if err := studentService.UpdateStudent(student); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
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
	studentRepo := repository.NewSQLRepository(config.GetConnection())
	studentService := service.NewStudentService(studentRepo)

	if err = studentService.DeleteStudent(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func StudentsHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		idStr := r.URL.Query().Get("id")
		if idStr != "" {
			GetStudentByIdHandler(w, r)
		} else {
			GetStudentsHandler(w, r)
		}
	case http.MethodPost:
		CreateStudentHandler(w, r)
	case http.MethodPut:
		UpdateStudentHandler(w, r)
	case http.MethodDelete:
		DeleteStudentHandler(w, r)
	default:
		http.Error(w, "Metodo no permitido", http.StatusBadRequest)
		return

	}

}
