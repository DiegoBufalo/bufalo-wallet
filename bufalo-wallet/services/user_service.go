package services

import (
	"bufalo-wallet-app/models"
	"bufalo-wallet-app/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *models.User) error {
	// Adicione lógica adicional de negócios aqui, se necessário
	return s.repo.Create(user)
}

func (s *UserService) GetUserByID(userID uint) (*models.User, error) {
	return s.repo.FindByID(userID)
}
