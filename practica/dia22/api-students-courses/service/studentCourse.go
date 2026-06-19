package service

import (
	"api-students-courses/domain"
	"api-students-courses/repository"
	"errors"
)

type StudentCourseService struct {
	studentCourseRepo repository.StudentCourseRepository
}

func NewStudentCourse(repo repository.StudentCourseRepository) *StudentCourseService {
	return &StudentCourseService{studentCourseRepo: repo}
}

func (s *StudentCourseService) AddStudentToCourse(student_id, course_id int64) error {

	if student_id <= 0 || course_id <= 0 {
		return errors.New("Id tiene que ser mayor igual que cero")
	}
	return s.studentCourseRepo.AddStudentToCourse(student_id, course_id)
}

func (s *StudentCourseService) GetStudentCourses(student_id int64) ([]domain.Course, error) {
	if student_id <= 0 {
		return nil, errors.New("Id tiene que ser mayor igual que cero")
	}
	return s.studentCourseRepo.GetStudentCourses(student_id)
}

func (s *StudentCourseService) GetStudentInCourses(course_id int64) ([]domain.Student, error) {
	if course_id <= 0 {
		return nil, errors.New("Id tiene que ser mayor igual que cero")
	}
	return s.studentCourseRepo.GetStudentInCourses(course_id)
}
