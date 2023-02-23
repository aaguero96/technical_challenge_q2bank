package utils

import "golang.org/x/crypto/bcrypt"

func ValidatePassword(hashedPassword, password string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false, err
	}
	return true, nil
}
