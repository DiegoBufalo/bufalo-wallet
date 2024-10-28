package services

import (
	"bufalo-wallet-app/models"
	"bufalo-wallet-app/repositories"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

// Adiciona uma nova categoria
func (s *CategoryService) CreateCategory(category *models.Category) error {
	return s.repo.Create(category)
}

// Retorna todas as categorias
func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	return s.repo.FindAll()
}

// Busca uma categoria específica por ID
func (s *CategoryService) GetCategoryByID(categoryID uint) (*models.Category, error) {
	return s.repo.FindByID(categoryID)
}
