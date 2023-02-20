package transactionService

import (
	"github.com/aaguero96/technical_challenge_q2bank/repository/transactionRepository"
)

type transactionService struct {
	transactionRepository transactionRepository.TransactionRepository
}

func NewTransactionService(transactionRepository transactionRepository.TransactionRepository) transactionService {
	return transactionService{
		transactionRepository: transactionRepository,
	}
}

func (ts transactionService) GetAll() ([]TransactionResponse, error) {
	transactions, err := ts.transactionRepository.GetAll()
	if err != nil {
		return nil, err
	}

	response := GetAllModelToResponse(transactions)

	return response, nil
}
