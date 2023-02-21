package transactionRepository

import (
	"errors"

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

func (tr transactionRepository) Deposit(payerID, payeeID, transactionID int, amount float64) error {
	// Create transaction
	tx := tr.db.Begin()

	// Verify if payer has credit
	var payer models.UserModel
	result := tx.First(&payer, payerID)
	if result.Error != nil {
		return result.Error
	}
	var payerWallet models.WalletModel
	result = tx.First(&payerWallet, payer.WalletID)
	if result.Error != nil {
		return result.Error
	}
	if payerWallet.Amount < amount {
		return errors.New("payer dont have enought money")
	}

	// Remove amount from payer wallet
	payerWallet.Amount -= amount
	result = tx.Model(&models.WalletModel{}).Where("wallet_id", payer.WalletID).Update("amount", payerWallet.Amount)
	if result.Error != nil {
		return result.Error
	}

	// Add amount in payee wallet
	var payee models.UserModel
	result = tx.First(&payee, payeeID)
	if result.Error != nil {
		return result.Error
	}
	var payeeWallet models.WalletModel
	result = tx.First(&payeeWallet, payee.WalletID)
	if result.Error != nil {
		return result.Error
	}
	payeeWallet.Amount += amount
	result = tx.Model(&models.WalletModel{}).Where("wallet_id", payee.WalletID).Update("amount", payeeWallet.Amount)
	if result.Error != nil {
		return result.Error
	}

	// Verify if transaction has status in progress
	var transaction models.TransactionModel
	result = tx.Where(&models.TransactionModel{Status: "in progress"}).First(&transaction, transactionID)
	if result.Error != nil {
		return result.Error
	}
	if transaction.Status != "in progress" {
		return result.Error
	}

	// Update transaction
	transaction.Status = "completed"
	result = tx.Save(&transaction)
	if result.Error != nil {
		return result.Error
	}

	// End Transaction
	tx.Commit()

	return nil
}

func (tr transactionRepository) DenyTransfer(transactionID int) error {
	var transaction models.TransactionModel
	result := tr.db.Where(&models.TransactionModel{Status: "in progress"}).First(&transaction, transactionID)
	if result.Error != nil {
		return result.Error
	}
	transaction.Status = "denied"
	result = tr.db.Save(&transaction)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (tr transactionRepository) CancelTransaction(transactionID int) error {
	var transaction models.TransactionModel
	result := tr.db.First(&transaction, transactionID)
	if result.Error != nil {
		return result.Error
	}

	if transaction.Status == "in progress" {
		transaction.Status = "canceled"
		result = tr.db.Save(&transaction)
		if result.Error != nil {
			return result.Error
		}
	}

	if transaction.Status == "completed" {
		transaction.Status = "cancel in progress"
		result = tr.db.Save(&transaction)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}
