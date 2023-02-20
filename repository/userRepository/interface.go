package userRepository

import "github.com/aaguero96/technical_challenge_q2bank/models"

type UserRepository interface {
	GetAll() ([]models.UserModel, error)
}
