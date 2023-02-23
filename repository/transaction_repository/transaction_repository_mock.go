package transaction_repository

import (
	"errors"

	"github.com/aaguero96/technical_challenge_q2bank/models"
	"gorm.io/gorm"
)

type transactionRepositoryMock struct {
	transactions []models.TransactionModel
}

func NewTransactionRepositoryMock() transactionRepositoryMock {
	transactions := []models.TransactionModel{
		{
			Model: gorm.Model{
				ID: 1,
			},
			PayerID: 1,
			PayeeID: 2,
			Amount:  100,
			Status:  "completed",
		},
	}

	return transactionRepositoryMock{transactions: transactions}
}

func (wrm transactionRepositoryMock) GetAll() ([]models.TransactionModel, error) {
	return wrm.transactions, nil
}

func (wrm transactionRepositoryMock) GetById(id int) (models.TransactionModel, error) {
	for _, transaction := range wrm.transactions {
		if int(transaction.ID) == id {
			return transaction, nil
		}
	}
	return models.TransactionModel{}, errors.New("transaction not found")
}

func (wrm transactionRepositoryMock) CreateTransaction(payerID, payeeID int, amount float64) (models.TransactionModel, error) {
	return models.TransactionModel{}, nil
}

func (wrm transactionRepositoryMock) Deposit(payerID, payeeID, transactionID int, amount float64) error {
	return nil
}

func (wrm transactionRepositoryMock) DenyTransfer(transactionID int) error {
	return nil
}

func (wrm transactionRepositoryMock) CancelTransaction(transactionID int) (string, error) {
	return "", nil
}

func (wrm transactionRepositoryMock) Return(payerID, payeeID, transactionID int, amount float64, status string) error {
	return nil
}
