package validator

type ValidatorExternalAPI interface {
	Validation() (bool, error)
}
