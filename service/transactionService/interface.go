package transactionService

type TransactionService interface {
	GetAll() ([]TransactionResponse, error)
	CreateTransaction(payerID, payeeID int, amount float64, payeerEmail string) (CreateTransactionResponse, error)
	Deposit(payerID, payeeID, transactionID int, amount float64) error
	DenyTransfer(transactionID int) error
	CancelTransaction(transactionID int, payeerEmail string) error
}
