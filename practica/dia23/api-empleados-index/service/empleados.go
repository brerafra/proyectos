package service

import (
	"errors"

	"main.go/domain"
	"main.go/repository"
)

type EmployeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) RegisterEmployee(Employee domain.Empleado) error {
	if Employee.Nombre == "" {
		return errors.New("Nombre requerido")
	}
	if Employee.Departamento == "" {
		return errors.New("Departamento requerido")
	}
	if Employee.Salario == 0 {
		return errors.New("Salario requerido")
	}

	if err := s.repo.Create(&Employee); err != nil {
		return err
	}
	return nil
}

func (s *EmployeeService) GetById(id int64) (*domain.Empleado, error) {
	if id <= 0 {
		return nil, errors.New("Id invalido")
	}
	return s.repo.GetById(id)
}

func (s *EmployeeService) GetAll(page, limit int) ([]domain.Empleado, int, error) {
	return s.repo.GetAll(page, limit)
}

func (s *EmployeeService) Update(e domain.Empleado) error {
	return s.repo.Update(&e)
}

func (s *EmployeeService) Delete(id int64) error {
	if id <= 0 {
		return errors.New("Id invalido")
	}
	return s.repo.Delete(id)
}
