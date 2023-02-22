package walletRepository

import (
	"errors"

	"github.com/aaguero96/technical_challenge_q2bank/models"
)

type walletRepositoryMock struct {
	wallets []models.WalletModel
}

func NewWalletRepositoryMock() walletRepositoryMock {
	wallets := []models.WalletModel{
		{WalletID: 1, Amount: 1000},
		{WalletID: 2, Amount: 2000},
	}

	return walletRepositoryMock{wallets: wallets}
}

func (wrm walletRepositoryMock) GetAll() ([]models.WalletModel, error) {
	return wrm.wallets, nil
}

func (wrm walletRepositoryMock) GetById(id int) (models.WalletModel, error) {
	for _, wallet := range wrm.wallets {
		if wallet.WalletID == id {
			return wallet, nil
		}
	}
	return models.WalletModel{}, errors.New("wallet not found")
}
