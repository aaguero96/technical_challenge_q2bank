package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateJWT(name, email, password string, userID int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = name
	claims["email"] = email
	claims["password"] = password
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Minute * 30)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, err
}
