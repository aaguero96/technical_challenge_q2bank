package walletService

import (
	"github.com/aaguero96/technical_challenge_q2bank/repository/walletRepository"
)

type walletService struct {
	walletRepository walletRepository.WalletRepository
}

func NewWalletService(wr walletRepository.WalletRepository) walletService {
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
