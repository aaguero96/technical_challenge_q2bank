package register_type_repository

import (
	"errors"

	"github.com/aaguero96/technical_challenge_q2bank/models"
)

type registerTypeRepositoryMock struct {
	registerTypes []models.RegisterTypeModel
}

func NewRegisterTypeRepositoryMock() registerTypeRepositoryMock {
	registerTypes := []models.RegisterTypeModel{
		{RegisterTypeID: 1, RegisterType: "CPF"},
		{RegisterTypeID: 2, RegisterType: "CNPJ"},
	}

	return registerTypeRepositoryMock{registerTypes: registerTypes}
}

func (wrm registerTypeRepositoryMock) GetAll() ([]models.RegisterTypeModel, error) {
	return wrm.registerTypes, nil
}

func (wrm registerTypeRepositoryMock) GetById(id int) (models.RegisterTypeModel, error) {
	for _, registerType := range wrm.registerTypes {
		if registerType.RegisterTypeID == id {
			return registerType, nil
		}
	}
	return models.RegisterTypeModel{}, errors.New("registerType not found")
}
