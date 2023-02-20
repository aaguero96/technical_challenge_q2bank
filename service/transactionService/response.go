package transactionService

import "github.com/aaguero96/technical_challenge_q2bank/models"

type TransactionResponse struct {
	PayerID int     `json:"payer_id"`
	PayeeID int     `json:"payee_id"`
	Amount  float64 `json:"amount"`
	Status  string  `json:"status"`
}

func GetAllModelToResponse(model []models.TransactionModel) []TransactionResponse {
	var response []TransactionResponse
	for _, value := range model {
		element := TransactionResponse{
			PayerID: value.PayerID,
			PayeeID: value.PayeeID,
			Amount:  value.Amount,
			Status:  value.Status,
		}
		response = append(response, element)
	}
	return response
}
