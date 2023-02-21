package userRepository

import "github.com/aaguero96/technical_challenge_q2bank/models"

type UserRepository interface {
	GetAll() ([]models.UserModel, error)
	GetById(id int) (models.UserModel, error)
	CreateUser(name, email, password string, registerNumber int64, registerTypeID, userTypeID int) (models.UserModel, error)
}
