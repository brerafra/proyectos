package service

import (
	"api-students-courses/domain"
	"api-students-courses/repository"
	"errors"
)

type StudentService struct {
	studentRepo repository.StudentRepository
}

func NewStudentService(repo repository.StudentRepository) *StudentService {
	return &StudentService{studentRepo: repo}
}

func (s *StudentService) RegisterStudent(name string) (*domain.Student, error) {
	if name == "" {
		return nil, errors.New("Nombre requerido")
	}

	student := &domain.Student{
		Name: name,
	}

	if err := s.studentRepo.Create(student); err != nil {
		return nil, err
	}
	return student, nil
}

func (s *StudentService) GetStudent(id int64) (*domain.Student, error) {
	if id <= 0 {
		return nil, errors.New("ID de estudiante invalido")
	}
	return s.studentRepo.GetById(id)
}

func (s *StudentService) GetStudents(page, limit int) ([]domain.Student, int, error) {
	return s.studentRepo.GetAll(page, limit)
}

func (s *StudentService) UpdateStudent(student domain.Student) error {
	return s.studentRepo.Update(&student)
}

func (s *StudentService) DeleteStudent(id int64) error {
	if id <= 0 {
		return errors.New("ID de estudiante invalido")
	}
	return s.studentRepo.Delete(id)
}
