package userTypeService

import "github.com/aaguero96/technical_challenge_q2bank/models"

type UserTypeResponse struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

func GetAllModelToResponse(model []models.UserTypeModel) []UserTypeResponse {
	var response []UserTypeResponse
	for _, value := range model {
		element := UserTypeResponse{
			ID:   value.UserTypeID,
			Type: value.UserType,
		}
		response = append(response, element)
	}
	return response
}

type GetByIdResponse struct {
	UserTypeID int    `json:"user_type_id"`
	UserType   string `json:"user_type"`
}

func GetByIdResponseModelToResponse(model models.UserTypeModel) GetByIdResponse {
	return GetByIdResponse{
		UserTypeID: model.UserTypeID,
		UserType:   model.UserType,
	}
}
