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

func (wr walletRepository) GetById(id int) (models.WalletModel, error) {
	var wallet models.WalletModel
	result := wr.db.First(&wallet, id)
	if result.Error != nil {
		return models.WalletModel{}, result.Error
	}
	return wallet, nil
}

func (wr *walletRepository) AddAmount(walletID int, increaseAmount float64) (models.WalletModel, error) {
	var wallet models.WalletModel
	result := wr.db.First(&wallet, walletID)
	if result.Error != nil {
		return models.WalletModel{}, result.Error
	}

	newAmount := wallet.Amount + increaseAmount
	result = wr.db.Model(&wallet).Update("amount", newAmount)
	if result.Error != nil {
		return models.WalletModel{}, result.Error
	}

	return wallet, nil
}
