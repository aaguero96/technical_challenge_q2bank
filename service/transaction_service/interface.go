package transaction_service

type TransactionService interface {
	GetAll() ([]TransactionResponse, error)
	CreateTransaction(payerID, payeeID int, amount float64, payeerEmail string) (CreateTransactionResponse, error)
	Deposit(payerID, payeeID, transactionID int, amount float64) error
	DenyTransfer(transactionID int) error
	CancelTransaction(transactionID int, payeerEmail string) error
	GetById(id int) (GetByIdResponse, error)
	Return(payerID, payeeID, transactionID int, amount float64, status string) error
}
