package transactionService

import (
	"errors"

	"github.com/aaguero96/technical_challenge_q2bank/events/producer"
	"github.com/aaguero96/technical_challenge_q2bank/repository/transactionRepository"
	"github.com/aaguero96/technical_challenge_q2bank/repository/userRepository"
	"github.com/aaguero96/technical_challenge_q2bank/repository/userTypeRepository"
	"github.com/aaguero96/technical_challenge_q2bank/repository/walletRepository"
)

type transactionService struct {
	transactionRepository transactionRepository.TransactionRepository
	userRepository        userRepository.UserRepository
	walletRepository      walletRepository.WalletRepository
	userTypeRepository    userTypeRepository.UserTypeRepository
}

func NewTransactionService(
	tr transactionRepository.TransactionRepository,
	ur userRepository.UserRepository,
	wr walletRepository.WalletRepository,
	utr userTypeRepository.UserTypeRepository,
) transactionService {
	return transactionService{
		transactionRepository: tr,
		userRepository:        ur,
		walletRepository:      wr,
		userTypeRepository:    utr,
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

func (ts transactionService) CreateTransaction(payerID, payeeID int, amount float64, payerEmail string) (CreateTransactionResponse, error) {
	// Verify if amount is more than zero
	if amount <= 0 {
		return CreateTransactionResponse{}, errors.New("amount has to be more than zero")
	}

	// Payer info
	payerData, err := ts.userRepository.GetById(payerID)
	if err != nil {
		return CreateTransactionResponse{}, err
	}

	// Verify if request payear is the same as logged
	if payerData.Email != payerEmail {
		return CreateTransactionResponse{}, errors.New("not authorized to do this transaction")
	}

	// Payee info
	payeeData, err := ts.userRepository.GetById(payeeID)
	if err != nil {
		return CreateTransactionResponse{}, err
	}

	// Verify if payeer and payee is not the same
	if payerData.ID == payeeData.ID {
		return CreateTransactionResponse{}, errors.New("payee and payer should be different")
	}

	// Payer wallet info
	walletPayerData, err := ts.walletRepository.GetById(payerID)
	if err != nil {
		return CreateTransactionResponse{}, err
	}

	// Verify if payer has enough money
	if walletPayerData.Amount < amount {
		return CreateTransactionResponse{}, errors.New("payer does not have enough money")
	}

	// Payer User Type info
	payerUserTypeData, err := ts.userTypeRepository.GetById(payeeData.UserTypeID)
	if err != nil {
		return CreateTransactionResponse{}, err
	}

	// Verify if payer is not storekeeper
	if payerUserTypeData.UserType == "storekeeper" {
		return CreateTransactionResponse{}, errors.New("storekeeper has not permission to transfer money")
	}

	// Create transaction
	transaction, err := ts.transactionRepository.CreateTransaction(payerID, payeeID, amount)
	if err != nil {
		return CreateTransactionResponse{}, err
	}

	response := CreateTransactionModelToResponse(transaction)

	// Send transaction to queue (for aprove)
	err = producer.ApprovingTransactionEvent(producer.TransactionEvent(response))
	if err != nil {
		return CreateTransactionResponse{}, err
	}

	return response, nil
}

func (ts transactionService) Deposit(payerID, payeeID, transactionID int, amount float64) error {
	err := ts.transactionRepository.Deposit(payerID, payeeID, transactionID, amount)
	if err != nil {
		return err
	}

	return nil
}

func (ts transactionService) DenyTransfer(transactionID int) error {
	err := ts.transactionRepository.DenyTransfer(transactionID)
	if err != nil {
		return err
	}

	return nil
}

func (ts transactionService) CancelTransaction(transactionID int, payeerEmail string) error {
	// Verify if payeerEmail is the same as payeer transaction
	transfer, err := ts.transactionRepository.GetById(transactionID)
	if err != nil {
		return err
	}
	payerData, err := ts.userRepository.GetById(transfer.PayerID)
	if err != nil {
		return err
	}
	if payerData.Email != payeerEmail {
		return errors.New("not authorized to do this transaction")
	}

	status, err := ts.transactionRepository.CancelTransaction(transactionID)
	if err != nil {
		return err
	}

	// Send transaction to queue (for aprove)
	err = producer.ApprovingTransactionEvent(producer.TransactionEvent{
		TransactionID: int(transfer.ID),
		PayerID:       transfer.PayerID,
		PayeeID:       transfer.PayeeID,
		Amount:        transfer.Amount,
		Status:        status,
	})
	if err != nil {
		return err
	}

	return nil
}

func (ts transactionService) GetById(id int) (GetByIdResponse, error) {
	user, err := ts.transactionRepository.GetById(id)
	if err != nil {
		return GetByIdResponse{}, err
	}

	response := GetByIdResponseModelToResponse(user)

	return response, nil
}
