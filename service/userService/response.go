package userService

import "github.com/aaguero96/technical_challenge_q2bank/models"

type UserResponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	WalletID int    `json:"wallet_id"`
}

func GetAllModelToResponse(model []models.UserModel) []UserResponse {
	var response []UserResponse
	for _, value := range model {
		element := UserResponse{
			Name:     value.Name,
			Email:    value.Email,
			WalletID: value.WalletID,
		}
		response = append(response, element)
	}
	return response
}
