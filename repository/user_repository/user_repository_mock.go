package user_repository

import (
	"errors"

	"github.com/aaguero96/technical_challenge_q2bank/models"
	"github.com/aaguero96/technical_challenge_q2bank/utils"
	"gorm.io/gorm"
)

type userRepositoryMock struct {
	users []models.UserModel
}

func NewUserRepositoryMock() userRepositoryMock {
	users := []models.UserModel{
		{
			Model: gorm.Model{
				ID: 1,
			},
			Name:           "name_1",
			RegisterNumber: 12345678900,
			RegisterTypeID: 1,
			Email:          "email1@test.com",
			WalletID:       1,
			UserTypeID:     1,
			Password:       "$2a$10$16yn5uLLPNhTcalc4XWDUuEgq.uiKywVEMa9SuM7KBN6iVGzMIYui",
		},
		{
			Model: gorm.Model{
				ID: 2,
			},
			Name:           "name_2",
			RegisterNumber: 12345678900001,
			RegisterTypeID: 2,
			Email:          "email2@test.com",
			WalletID:       2,
			UserTypeID:     2,
			Password:       "$2a$10$WGJoxCGpj.rAm3CCU3z8pOXdgm4VCZhsOkrGIj7f2O2/DLwYupXKy",
		},
	}

	return userRepositoryMock{users: users}
}

func (utrm userRepositoryMock) GetAll() ([]models.UserModel, error) {
	return utrm.users, nil
}

func (utrm userRepositoryMock) GetById(id int) (models.UserModel, error) {
	for _, user := range utrm.users {
		if int(user.ID) == id {
			return user, nil
		}
	}
	return models.UserModel{}, errors.New("user not found")
}

func (utrm *userRepositoryMock) CreateUser(name, email, password string, registerNumber int64, registerTypeID, userTypeID int) (models.UserModel, error) {
	nextId := len(utrm.users) + 1
	utrm.users = append(utrm.users, models.UserModel{
		Model: gorm.Model{
			ID: uint(nextId),
		},
		Name:           name,
		RegisterNumber: registerNumber,
		RegisterTypeID: registerTypeID,
		Email:          email,
		WalletID:       nextId,
		UserTypeID:     userTypeID,
		Password:       password,
	})

	return utrm.users[nextId-1], nil
}

func (utrm userRepositoryMock) LoginUser(email, password string) (models.UserModel, error) {
	for _, user := range utrm.users {
		vaidatePassword, _ := utils.ValidatePassword(user.Password, password)
		validatorEmail := user.Email == email
		validate := vaidatePassword && validatorEmail
		if !validate {
			return user, nil
		}
	}
	return models.UserModel{}, errors.New("unauthorized")
}
