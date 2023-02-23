package wallet_repository

import "github.com/aaguero96/technical_challenge_q2bank/models"

type WalletRepository interface {
	GetAll() ([]models.WalletModel, error)
	GetById(id int) (models.WalletModel, error)
	AddAmount(walletID int, increaseAmount float64) (models.WalletModel, error)
}
