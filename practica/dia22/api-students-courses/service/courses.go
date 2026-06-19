package service

import (
	"api-students-courses/domain"
	"api-students-courses/repository"
	"errors"
)

type CourseService struct {
	courseRepo repository.CourseRepository
}

func NewCourseService(repo repository.CourseRepository) *CourseService {
	return &CourseService{courseRepo: repo}
}

func (s *CourseService) RegisterCourse(name string) (*domain.Course, error) {
	if name == "" {
		return nil, errors.New("Nombre requerido")
	}

	course := &domain.Course{
		Name: name,
	}

	if err := s.courseRepo.Create(course); err != nil {
		return nil, err
	}
	return course, nil
}

func (s *CourseService) GetCourse(id int64) (*domain.Course, error) {
	if id <= 0 {
		return nil, errors.New("ID de curso invalido")
	}
	return s.courseRepo.GetById(id)
}

func (s *CourseService) GetCourses(page, limit int) ([]domain.Course, int, error) {
	return s.courseRepo.GetAll(page, limit)
}

func (s *CourseService) UpdateCourse(course domain.Course) error {
	return s.courseRepo.Update(&course)
}

func (s *CourseService) DeleteCourse(id int64) error {
	if id <= 0 {
		return errors.New("ID de curso invalido")
	}
	return s.courseRepo.Delete(id)
}
