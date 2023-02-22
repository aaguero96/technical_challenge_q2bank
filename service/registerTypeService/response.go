package registerTypeService

import "github.com/aaguero96/technical_challenge_q2bank/models"

type RegisterTypeResponse struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

func GetAllModelToResponse(model []models.RegisterTypeModel) []RegisterTypeResponse {
	var response []RegisterTypeResponse
	for _, value := range model {
		element := RegisterTypeResponse{
			ID:   value.RegisterTypeID,
			Type: value.RegisterType,
		}
		response = append(response, element)
	}
	return response
}
