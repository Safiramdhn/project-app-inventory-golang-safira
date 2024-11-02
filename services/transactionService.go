package services

import (
	"errors"
	"time"

	"github.com/Safiramdhn/project-app-inventory-golang-safira/models"
	"github.com/Safiramdhn/project-app-inventory-golang-safira/repositories"
)

type Transaction struct {
	RepoTransaction repositories.TransactionRepositoryDB
}

func NewTransaction(repo repositories.TransactionRepositoryDB) *Transaction {
	return &Transaction{RepoTransaction: repo}
}

func (service *Transaction) CreateTransaction(transaction *models.Transaction) error {
	if transaction.Quantity < 0 {
		return errors.New("invalid quantity")
	}
	if transaction.Description == "" {
		return errors.New("invalid description")
	}
	if transaction.AddedBy <= 0 {
		return errors.New("invalid added by")
	}
	if transaction.Item.ID <= 0 {
		return errors.New("invalid item ID")
	}

	transaction.Timestamp = time.Now()
	return service.RepoTransaction.Add(transaction)
}

func (service *Transaction) GetTransactionByID(id int) (*models.Transaction, error) {
	if id <= 0 {
		return nil, errors.New("invalid transaction ID")
	}

	return service.RepoTransaction.GetByID(id)
}

// func (service *Transaction) UpdateTransaction(transaction *models.Transaction) error {
// 	if transaction.ID <= 0 {
// 		return errors.New("invalid transaction ID")
// 	}

// 	return service.RepoTransaction.Update(transaction)
// }

// func (service *Transaction) DeleteTransaction(id int) error {
// 	if id <= 0 {
// 		return errors.New("invalid transaction ID")
// 	}

// 	return service.RepoTransaction.Delete(id)
// }

func (service *Transaction) GetAllTransactions(pagination models.Pagination) ([]models.Transaction, error) {
	var limit, offset int
	if pagination.PerPage != 0 {
		limit = pagination.PerPage
		offset = (pagination.Page - 1) * pagination.PerPage
	}
	return service.RepoTransaction.GetAll(limit, offset)
}

func (service *Transaction) GetAllTransactionsWithFilter(transaction models.Transaction) ([]models.Transaction, error) {
	var limit, offset int
	if transaction.Pagination.PerPage != 0 {
		limit = transaction.Pagination.PerPage
		offset = (transaction.Pagination.Page - 1) * transaction.Pagination.PerPage
	}
	return service.RepoTransaction.GetAllWithFilter(transaction, limit, offset)
}
