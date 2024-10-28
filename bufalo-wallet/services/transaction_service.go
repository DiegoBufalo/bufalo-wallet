package services

import (
	"bufalo-wallet-app/models"
	"bufalo-wallet-app/repositories"
)

type TransactionService struct {
	repo *repositories.TransactionRepository
}

func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) CreateTransaction(transaction *models.Transaction) error {
	// Validação ou lógica adicional, se necessário
	return s.repo.Create(transaction)
}

func (s *TransactionService) GetUserTransactions(userID uint) ([]models.Transaction, error) {
	return s.repo.FindByUserID(userID)
}

func (s *TransactionService) UpdateTransaction(transaction *models.Transaction) error {
	existingTransaction, err := s.repo.FindByID(transaction.ID)
	if err != nil {
		return err
	}

	// Exemplo de atualização de dados (pode ser adaptado)
	existingTransaction.Amount = transaction.Amount
	existingTransaction.Date = transaction.Date
	existingTransaction.CategoryID = transaction.CategoryID
	existingTransaction.Type = transaction.Type
	existingTransaction.Recurring = transaction.Recurring
	existingTransaction.PeriodDays = transaction.PeriodDays

	return s.repo.Update(existingTransaction)
}

func (s *TransactionService) DeleteTransaction(transactionID uint) error {
	return s.repo.Delete(transactionID)
}
