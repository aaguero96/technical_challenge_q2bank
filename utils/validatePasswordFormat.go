package utils

import "unicode"

func ValidatePasswordFormat(password string) (string, bool) {
	type Validate struct {
		validate bool
		message  string
	}

	var validations = make(map[string]Validate)
	validations["hasMininumLen"] = Validate{validate: false, message: "password has no minimum characters length"}
	validations["hasUpperChar"] = Validate{validate: false, message: "password has no upper case characters"}
	validations["hasLowerChar"] = Validate{validate: false, message: "password has no lower case characters"}
	validations["hasNumberChar"] = Validate{validate: false, message: "password has no numbers characters"}
	validations["hasSpecialChar"] = Validate{validate: false, message: "password has no apecial characters"}

	if len(password) >= 5 {
		validations["hasMininumLen"] = Validate{validate: true, message: validations["hasMininumLen"].message}
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			validations["hasUpperChar"] = Validate{validate: true, message: validations["hasUpperChar"].message}
		case unicode.IsLower(char):
			validations["hasLowerChar"] = Validate{validate: true, message: validations["hasLowerChar"].message}
		case unicode.IsNumber(char):
			validations["hasNumberChar"] = Validate{validate: true, message: validations["hasNumberChar"].message}
		case unicode.IsPunct(char):
			validations["hasSpecialChar"] = Validate{validate: true, message: validations["hasSpecialChar"].message}
		}
	}

	for _, value := range validations {
		if !value.validate {
			return value.message, false
		}
	}
	return "", true
}
