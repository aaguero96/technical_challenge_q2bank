package registerTypeService

import "github.com/aaguero96/technical_challenge_q2bank/models"

type RegiterTypeResponse struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

func GetAllModelToResponse(model []models.RegisterTypeModel) []RegiterTypeResponse {
	var response []RegiterTypeResponse
	for _, value := range model {
		element := RegiterTypeResponse{
			ID:   value.RegisterTypeID,
			Type: value.RegisterType,
		}
		response = append(response, element)
	}
	return response
}
