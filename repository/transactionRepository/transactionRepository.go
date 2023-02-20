package transactionRepository

import (
	"github.com/aaguero96/technical_challenge_q2bank/models"
	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) transactionRepository {
	return transactionRepository{
		db: db,
	}
}

func (tr transactionRepository) GetAll() ([]models.TransactionModel, error) {
	var transactions []models.TransactionModel
	result := tr.db.Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return transactions, nil
}

func (tr transactionRepository) CreateTransaction(payerID, payeeID int, amount float64) (models.TransactionModel, error) {
	transaction := models.TransactionModel{
		PayerID: payerID,
		PayeeID: payeeID,
		Amount:  amount,
		Status:  "in progress",
	}
	result := tr.db.Create(&transaction)
	if result.Error != nil {
		return models.TransactionModel{}, result.Error
	}
	return transaction, nil
}
