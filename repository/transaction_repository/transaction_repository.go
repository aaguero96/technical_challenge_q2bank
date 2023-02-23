package transaction_repository

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

func (tr transactionRepository) CancelTransaction(transactionID int) (string, error) {
	var transaction models.TransactionModel
	result := tr.db.First(&transaction, transactionID)
	if result.Error != nil {
		return "", result.Error
	}

	// Verify if payee has necessary amount to retun
	var payeeData models.UserModel
	result = tr.db.First(&payeeData, transaction.PayeeID)
	if result.Error != nil {
		return "", result.Error
	}
	var payeeWalletData models.WalletModel
	result = tr.db.First(&payeeWalletData, payeeData.WalletID)
	if result.Error != nil {
		return "", result.Error
	}
	if payeeWalletData.Amount < transaction.Amount {
		return "", errors.New("payee has not necessary amount to return please contact phone... to more info")
	}

	if transaction.Status == "in progress" {
		transaction.Status = "canceled"
		result = tr.db.Save(&transaction)
		if result.Error != nil {
			return "", result.Error
		}
	}

	if transaction.Status == "completed" {
		transaction.Status = "cancel in progress"
		result = tr.db.Save(&transaction)
		if result.Error != nil {
			return "", result.Error
		}
	}
	return transaction.Status, nil
}

func (tr transactionRepository) GetById(id int) (models.TransactionModel, error) {
	var transaction models.TransactionModel
	result := tr.db.First(&transaction, id)
	if result.Error != nil {
		return models.TransactionModel{}, result.Error
	}
	return transaction, nil
}

func (tr transactionRepository) Return(payerID, payeeID, transactionID int, amount float64, status string) error {
	// If is canceled nothing happen
	if status == "canceled" {
		return nil
	}

	// // If is in progress amount will be returned
	// Create transaction
	tx := tr.db.Begin()

	// Verify if payee has credit
	var payee models.UserModel
	result := tx.First(&payee, payeeID)
	if result.Error != nil {
		return result.Error
	}
	var payeeWallet models.WalletModel
	result = tx.First(&payeeWallet, payee.WalletID)
	if result.Error != nil {
		return result.Error
	}
	if payeeWallet.Amount < amount {
		return errors.New("payee dont have enought money to return")
	}

	// Remove amount from payer wallet
	payeeWallet.Amount -= amount
	result = tx.Model(&models.WalletModel{}).Where("wallet_id", payee.WalletID).Update("amount", payeeWallet.Amount)
	if result.Error != nil {
		return result.Error
	}

	// Add amount in payer wallet
	var payer models.UserModel
	result = tx.First(&payer, payerID)
	if result.Error != nil {
		return result.Error
	}
	var payerWallet models.WalletModel
	result = tx.First(&payerWallet, payer.WalletID)
	if result.Error != nil {
		return result.Error
	}
	payerWallet.Amount += amount
	result = tx.Model(&models.WalletModel{}).Where("wallet_id", payer.WalletID).Update("amount", payerWallet.Amount)
	if result.Error != nil {
		return result.Error
	}

	// Update transaction
	var transaction models.TransactionModel
	result = tx.Where(&models.TransactionModel{Status: "cancel in progress"}).First(&transaction, transactionID)
	if result.Error != nil {
		return result.Error
	}
	transaction.Status = "canceled"
	result = tx.Save(&transaction)
	if result.Error != nil {
		return result.Error
	}

	// End Transaction
	tx.Commit()

	return nil
}
