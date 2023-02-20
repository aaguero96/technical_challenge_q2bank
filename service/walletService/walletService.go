package walletService

import (
	"github.com/aaguero96/technical_challenge_q2bank/repository/walletRepository"
)

type walletService struct {
	walletRepository walletRepository.WalletRepository
}

func NewWalletService(walletRepository walletRepository.WalletRepository) walletService {
	return walletService{
		walletRepository: walletRepository,
	}
}

func (ws walletService) GetAll() ([]WalletResponse, error) {
	wallets, err := ws.walletRepository.GetAll()
	if err != nil {
		return nil, err
	}

	response := GetAllModelToResponse(wallets)

	return response, nil
}