package userTypeRepository

import "github.com/aaguero96/technical_challenge_q2bank/models"

type UserTypeRepository interface {
	GetAll() ([]models.UserTypeModel, error)
	GetById(id int) (models.UserTypeModel, error)
}
