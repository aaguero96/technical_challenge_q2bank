package registerTypeRepository

import "github.com/aaguero96/technical_challenge_q2bank/models"

type RegisterTypeRepository interface {
	GetAll() ([]models.RegisterTypeModel, error)
}
