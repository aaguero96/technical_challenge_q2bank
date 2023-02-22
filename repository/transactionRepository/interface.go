package transactionRepository

import "github.com/aaguero96/technical_challenge_q2bank/models"

type TransactionRepository interface {
	GetAll() ([]models.TransactionModel, error)
	CreateTransaction(payerID, payeeID int, amount float64) (models.TransactionModel, error)
	Deposit(payerID, payeeID, transactionID int, amount float64) error
	DenyTransfer(transactionID int) error
	CancelTransaction(transactionID int) (string, error)
	GetById(id int) (models.TransactionModel, error)
	Return(payerID, payeeID, transactionID int, amount float64, status string) error
}
