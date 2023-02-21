package externalValidatorService

type ExternalValidatorService interface {
	Validator() (bool, error)
}
