package external_validator_service

type ExternalValidatorService interface {
	Validator() (bool, error)
}
