package transactionService

type TransactionService interface {
	GetAll() ([]TransactionResponse, error)
	CreateTransaction(payerID, payeeID int, amount float64) (CreateTransactionResponse, error)
	Deposit(payerID, payeeID, transactionID int, amount float64) error
}
