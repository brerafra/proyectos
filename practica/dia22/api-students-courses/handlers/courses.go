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

type PaginatedCourseResponse struct {
	Data       []domain.Course `json:"data"`
	TotalRows  int             `json:"total_rows"`
	TotalPages int             `json:"total_pages"`
	Page       int             `json:"page"`
	Limit      int             `json:"limit"`
}

func CreateCourseHandler(w http.ResponseWriter, r *http.Request) {
	var course domain.Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	courseRepo := repository.NewSQLCourseRepository(config.GetConnection())
	CourseService := service.NewCourseService(courseRepo)

	if _, err := CourseService.RegisterCourse(course.Name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetCourseByIdHandler(w http.ResponseWriter, r *http.Request) {
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

	courseRepo := repository.NewSQLCourseRepository(config.GetConnection())
	CourseService := service.NewCourseService(courseRepo)

	course, err := CourseService.GetCourse(id)
	if err != nil {
		http.Error(w, "No se encontro usuario", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)
}

func GetCoursesHandler(w http.ResponseWriter, r *http.Request) {
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

	courseRepo := repository.NewSQLCourseRepository(config.GetConnection())
	CourseService := service.NewCourseService(courseRepo)

	courses, totalRows, err := CourseService.GetCourses(page, limit)

	totalPages := (totalRows + limit - 1) / limit
	response := PaginatedCourseResponse{
		Data:       courses,
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Page:       page,
		Limit:      limit,
	}

	w.Header().Set("Content-Type", "applitation/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateCourseHandler(w http.ResponseWriter, r *http.Request) {
	var course domain.Course

	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	courseRepo := repository.NewSQLCourseRepository(config.GetConnection())
	courseService := service.NewCourseService(courseRepo)

	if err := courseService.UpdateCourse(course); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteCourseHandler(w http.ResponseWriter, r *http.Request) {
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
	courseRepo := repository.NewSQLCourseRepository(config.GetConnection())
	courseService := service.NewCourseService(courseRepo)

	if err = courseService.DeleteCourse(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func CoursesHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		idStr := r.URL.Query().Get("id")
		if idStr != "" {
			GetCourseByIdHandler(w, r)
		} else {
			GetCoursesHandler(w, r)
		}
	case http.MethodPost:
		CreateCourseHandler(w, r)
	case http.MethodPut:
		UpdateCourseHandler(w, r)
	case http.MethodDelete:
		DeleteCourseHandler(w, r)
	default:
		http.Error(w, "Metodo no permitido", http.StatusBadRequest)
		return

	}

}
