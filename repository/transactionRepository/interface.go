package transactionRepository

import "github.com/aaguero96/technical_challenge_q2bank/models"

type TransactionRepository interface {
	GetAll() ([]models.TransactionModel, error)
}
