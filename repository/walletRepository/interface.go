package walletRepository

import "github.com/aaguero96/technical_challenge_q2bank/models"

type WalletRepository interface {
	GetAll() ([]models.WalletModel, error)
	GetById(id int) (models.WalletModel, error)
}
