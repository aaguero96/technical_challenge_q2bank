package transactionService

import (
	"errors"

	"github.com/aaguero96/technical_challenge_q2bank/events/producer"
	"github.com/aaguero96/technical_challenge_q2bank/repository/transactionRepository"
	"github.com/aaguero96/technical_challenge_q2bank/repository/userRepository"
	"github.com/aaguero96/technical_challenge_q2bank/repository/walletRepository"
)

type transactionService struct {
	transactionRepository transactionRepository.TransactionRepository
	userRepository        userRepository.UserRepository
	walletRepository      walletRepository.WalletRepository
}

func NewTransactionService(
	tr transactionRepository.TransactionRepository,
	ur userRepository.UserRepository,
	wr walletRepository.WalletRepository,
) transactionService {
	return transactionService{
		transactionRepository: tr,
		userRepository:        ur,
		walletRepository:      wr,
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

func (ts transactionService) CreateTransaction(payerID, payeeID int, amount float64) (CreateTransactionResponse, error) {
	payerData, err := ts.userRepository.GetById(payerID)
	if err != nil {
		return CreateTransactionResponse{}, err
	}

	payeeData, err := ts.userRepository.GetById(payeeID)
	if err != nil {
		return CreateTransactionResponse{}, err
	}

	if payerData.ID == payeeData.ID {
		return CreateTransactionResponse{}, errors.New("payee and payer should be different")
	}

	walletPayerData, err := ts.walletRepository.GetById(payerID)
	if err != nil {
		return CreateTransactionResponse{}, err
	}

	if walletPayerData.Amount < amount {
		return CreateTransactionResponse{}, errors.New("payer dont have enough money")
	}

	transaction, err := ts.transactionRepository.CreateTransaction(payerID, payeeID, amount)
	if err != nil {
		return CreateTransactionResponse{}, err
	}

	response := CreateTransactionModelToResponse(transaction)

	err = producer.ApprovingTransactionEvent(producer.TransactionEvent(response))
	if err != nil {
		return CreateTransactionResponse{}, err
	}

	return response, nil
}
