package utils

import (
	"errors"
	"fmt"
)

func ValidateRegisterNumber(registerNumber int64, registerType string) error {
	var validRegisterNumber bool = true
	switch registerType {
	case "CPF":
		validRegisterNumber = len(fmt.Sprintf("%v", registerNumber)) == 11
	case "CNPJ":
		validRegisterNumber = len(fmt.Sprintf("%v", registerNumber)) == 14
	default:
		return errors.New("register_type is incorrect")
	}
	if !validRegisterNumber {
		return errors.New("register number is invalid, that was considering the register type passed")
	}
	return nil
}
