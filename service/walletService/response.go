package walletService

import "github.com/aaguero96/technical_challenge_q2bank/models"

type WalletResponse struct {
	ID     int     `json:"id"`
	Amount float64 `json:"amount"`
}

func GetAllModelToResponse(model []models.WalletModel) []WalletResponse {
	var response []WalletResponse
	for _, value := range model {
		element := WalletResponse{
			ID:     value.WalletID,
			Amount: value.Amount,
		}
		response = append(response, element)
	}
	return response
}

type GetByIdResponse struct {
	WalletID int     `json:"wallet_id"`
	Amount   float64 `json:"amount"`
}

func GetByIdModelToResponse(model models.WalletModel) GetByIdResponse {
	return GetByIdResponse{
		WalletID: model.WalletID,
		Amount:   model.Amount,
	}
}
