package validator

type validatorExternalAPIMock struct {
	authorization bool
}

func NewValidatorExternalAPIMock(authorization bool) validatorExternalAPIMock {
	return validatorExternalAPIMock{
		authorization: authorization,
	}
}

func (veam validatorExternalAPIMock) Validation() (bool, error) {
	return veam.authorization, nil
}
