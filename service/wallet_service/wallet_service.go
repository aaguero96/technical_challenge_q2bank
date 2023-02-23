package wallet_service

import (
	"github.com/aaguero96/technical_challenge_q2bank/repository/wallet_repository"
)

type walletService struct {
	walletRepository wallet_repository.WalletRepository
}

func NewWalletService(wr wallet_repository.WalletRepository) walletService {
	return walletService{
		walletRepository: wr,
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

func (ws walletService) GetById(id int) (GetByIdResponse, error) {
	wallet, err := ws.walletRepository.GetById(id)
	if err != nil {
		return GetByIdResponse{}, err
	}

	response := GetByIdModelToResponse(wallet)

	return response, nil
}

func (ws walletService) AddAmount(walletID int, increaseAmount float64) (AddAmountResponse, error) {
	wallet, err := ws.walletRepository.AddAmount(walletID, increaseAmount)
	if err != nil {
		return AddAmountResponse{}, err
	}

	response := AddAmountModelToResponse(wallet)

	return response, nil
}
