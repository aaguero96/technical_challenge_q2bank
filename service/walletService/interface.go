package walletService

import "github.com/aaguero96/technical_challenge_q2bank/models"

type WalletService interface {
	GetAll() ([]WalletResponse, error)
	GetById(id int) (GetByIdResponse, error)
	AddAmount(walletID int, increaseAmount float64) (models.WalletModel, error)
}
