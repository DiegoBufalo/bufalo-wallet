package repositories

import (
	"bufalo-wallet-app/models"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Create(transaction *models.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *TransactionRepository) FindByUserID(userID uint) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Where("user_id = ?", userID).Find(&transactions).Error
	return transactions, err
}

func (r *TransactionRepository) FindByID(transactionID uint) (*models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.First(&transaction, transactionID).Error
	return &transaction, err
}

func (r *TransactionRepository) Update(transaction *models.Transaction) error {
	return r.db.Save(transaction).Error
}

func (r *TransactionRepository) Delete(transactionID uint) error {
	return r.db.Delete(&models.Transaction{}, transactionID).Error
}
