package service

import (
	"api-students-courses/domain"
	"api-students-courses/repository"
	"errors"
)

type TeacherService struct {
	repo repository.TeacherRepository
}

func NewTeacherService(repo repository.TeacherRepository) *TeacherService {
	return &TeacherService{repo: repo}
}

func (s *TeacherService) RegisterTeacher(name string, shift int) error {
	if name == "" {
		return errors.New("Nombre requerido")
	}

	if shift <= 0 {
		return errors.New("turno requerido")
	}

	teacher := &domain.Teacher{
		Name:  name,
		Shift: shift,
	}

	if err := s.repo.Create(teacher); err != nil {
		return err
	}
	return nil
}

func (s *TeacherService) GetTeacher(id int64) (*domain.Teacher, error) {
	if id <= 0 {
		return nil, errors.New("ID invalido")
	}
	return s.repo.GetById(id)
}

func (s *TeacherService) GetTeachers(page, limit int) ([]domain.Teacher, int, error) {
	return s.repo.GetAll(page, limit)
}

func (s *TeacherService) UpdateTeacher(teacher domain.Teacher) error {
	return s.repo.Update(&teacher)
}

func (s *TeacherService) DeleteTeacher(id int64) error {
	if id <= 0 {
		return errors.New("ID invalido ")
	}
	return s.repo.Delete(id)
}
