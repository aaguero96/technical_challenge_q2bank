package external_validator_service

import (
	"github.com/aaguero96/technical_challenge_q2bank/external_API/validator"
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
