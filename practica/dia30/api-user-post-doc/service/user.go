package service

import (
	"errors"
	"main/domain"
	"main/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(u domain.User) error {
	if u.Name == "" {
		return errors.New("Nombre requerido")
	}

	if u.Email == "" {
		return errors.New("Email requerido")
	}

	if u.Card <= 0 {
		return errors.New("tarjeta requerida")
	}

	user := &domain.User{
		Name:     u.Name,
		Email:    u.Email,
		Card:     u.Card,
		IsActive: true,
		IsAdmin:  false,
	}

	if err := s.repo.Create(user); err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetUser(id int64) (*domain.User, error) {
	if id <= 0 {
		return nil, errors.New("Id invalida")
	}
	return s.repo.GetById(id)
}

func (s *UserService) GetUsers(page, limit int) ([]domain.User, int, error) {
	return s.repo.GetAll(page, limit)
}

func (s *UserService) UpdateUser(u domain.User) error {
	return s.repo.Update(&u)
}

func (s *UserService) DeleteUser(id int64) error {
	if id <= 0 {
		return errors.New("Id invalida")
	}
	return s.repo.Delete(id)
}
