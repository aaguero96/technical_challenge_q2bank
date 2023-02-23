package user_type_repository

import (
	"errors"

	"github.com/aaguero96/technical_challenge_q2bank/models"
)

type userTypeRepositoryMock struct {
	userTypes []models.UserTypeModel
}

func NewUserTypeRepositoryMock() userTypeRepositoryMock {
	userTypes := []models.UserTypeModel{
		{UserTypeID: 1, UserType: "common"},
		{UserTypeID: 2, UserType: "storekeeper"},
	}

	return userTypeRepositoryMock{userTypes: userTypes}
}

func (utrm userTypeRepositoryMock) GetAll() ([]models.UserTypeModel, error) {
	return utrm.userTypes, nil
}

func (utrm userTypeRepositoryMock) GetById(id int) (models.UserTypeModel, error) {
	for _, wallet := range utrm.userTypes {
		if wallet.UserTypeID == id {
			return wallet, nil
		}
	}
	return models.UserTypeModel{}, errors.New("user type not found")
}
