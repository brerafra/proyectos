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

func AddStudentToCourseHandler(w http.ResponseWriter, r *http.Request) {
	var s domain.StudentCourse
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	studentCourseRepo := repository.NewSQLStudentCourseRepository(config.GetConnection())
	studentCourseService := service.NewStudentCourse(studentCourseRepo)

	if err := studentCourseService.AddStudentToCourse(s.StudentId, s.CourseId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func GetCourseOfAStudent(w http.ResponseWriter, r *http.Request) {
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

	studentCourseRepo := repository.NewSQLStudentCourseRepository(config.GetConnection())
	studentCourseService := service.NewStudentCourse(studentCourseRepo)

	students, err := studentCourseService.GetStudentCourses(id)
	if err != nil {
		http.Error(w, "No se encontro estudiantes", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)

}

func GetStudentInCoursesHandler(w http.ResponseWriter, r *http.Request) {
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

	studentCourseRepo := repository.NewSQLStudentCourseRepository(config.GetConnection())
	studentCourseService := service.NewStudentCourse(studentCourseRepo)

	students, err := studentCourseService.GetStudentInCourses(id)
	if err != nil {
		http.Error(w, "No se encontro estudiantes", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func CourseStudentsHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetCourseOfAStudent(w, r)
	default:
		http.Error(w, "Metodo no permitido", http.StatusBadRequest)
		return

	}

}

func StudentCoursesHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetStudentInCoursesHandler(w, r)
	case http.MethodPost:
		AddStudentToCourseHandler(w, r)
	default:
		http.Error(w, "Metodo no permitido", http.StatusBadRequest)
		return

	}

}
