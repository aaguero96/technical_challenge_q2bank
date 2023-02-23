package transaction_handler

type CreateTransactionRequest struct {
	PayerID int     `json:"payer_id"`
	PayeeID int     `json:"payee_id"`
	Amount  float64 `json:"amount"`
}
