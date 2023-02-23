package user_service

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

type GetByIdResponse struct {
	Name           string `json:"name"`
	RegisterNumber int64  `json:"register_number"`
	RegisterTypeID int    `json:"register_type_id"`
	Email          string `json:"email"`
	WalletID       int    `json:"wallet_id"`
	UserTypeID     int    `json:"user_type_id"`
}

func GetByIdResponseModelToResponse(model models.UserModel) GetByIdResponse {
	return GetByIdResponse{
		Name:           model.Name,
		RegisterNumber: model.RegisterNumber,
		RegisterTypeID: model.RegisterTypeID,
		Email:          model.Email,
		WalletID:       model.WalletID,
		UserTypeID:     model.UserTypeID,
	}
}

type CreateUserResponse struct {
	Token      string `json:"token"`
	ExpiringIn string `json:"expiring_in"`
}

type LoginUserResponse struct {
	Token      string `json:"token"`
	ExpiringIn string `json:"expiring_in"`
}
