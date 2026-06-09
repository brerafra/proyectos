package service

import (
	"errors"
	"main/domain"
	"main/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{userRepo: repo}
}

func (s *UserService) RegisterUSer(name, email string) (*domain.User, error) {
	if name == "" || email == "" {
		return nil, errors.New("nombre y email son requeridos")
	}

	user := &domain.User{
		Name:  name,
		Email: email,
	}
	err := s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUser(id int) (*domain.User, error) {
	if id <= 0 {
		return nil, errors.New("ID invalido")
	}
	return s.userRepo.GetById(id)
}
