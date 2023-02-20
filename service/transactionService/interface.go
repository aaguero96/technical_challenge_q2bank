package transactionService

type TransactionService interface {
	GetAll() ([]TransactionResponse, error)
}
