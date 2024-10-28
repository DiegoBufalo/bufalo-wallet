package repositories

import (
	"bufalo-wallet-app/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

// Cria uma nova categoria
func (r *CategoryRepository) Create(category *models.Category) error {
	return r.db.Create(category).Error
}

// Busca todas as categorias
func (r *CategoryRepository) FindAll() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

// Busca uma categoria pelo ID
func (r *CategoryRepository) FindByID(categoryID uint) (*models.Category, error) {
	var category models.Category
	err := r.db.First(&category, categoryID).Error
	return &category, err
}
