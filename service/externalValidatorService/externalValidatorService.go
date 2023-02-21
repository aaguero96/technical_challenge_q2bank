package externalValidatorService

import (
	"github.com/aaguero96/technical_challenge_q2bank/externalAPI/validator"
)

type externalValidatorService struct {
	externalValidatorAPI validator.ValidatorExternalAPI
}

func NewExternalValidatorService(externalValidatorAPI validator.ValidatorExternalAPI) externalValidatorService {
	return externalValidatorService{
		externalValidatorAPI: externalValidatorAPI,
	}
}

func (evs externalValidatorService) Validator() (bool, error) {
	response, err := evs.externalValidatorAPI.Validation()
	if err != nil {
		return false, err
	}
	return response, nil
}
