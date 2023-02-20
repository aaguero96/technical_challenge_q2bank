package walletRepository

import (
	"github.com/aaguero96/technical_challenge_q2bank/models"
	"gorm.io/gorm"
)

type walletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) walletRepository {
	return walletRepository{
		db: db,
	}
}

func (wr walletRepository) GetAll() ([]models.WalletModel, error) {
	var wallets []models.WalletModel
	result := wr.db.Find(&wallets)
	if result.Error != nil {
		return nil, result.Error
	}
	return wallets, nil
}
