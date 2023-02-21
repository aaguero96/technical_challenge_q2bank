package userService

import (
	"fmt"

	"github.com/aaguero96/technical_challenge_q2bank/repository/userRepository"
)

type userService struct {
	userRepository userRepository.UserRepository
}

func NewUserService(ur userRepository.UserRepository) userService {
	return userService{
		userRepository: ur,
	}
}

func (us userService) GetAll() ([]UserResponse, error) {
	users, err := us.userRepository.GetAll()
	if err != nil {
		return nil, err
	}

	response := GetAllModelToResponse(users)

	return response, nil
}

func (us userService) GetById(id int) (GetByIdResponse, error) {
	user, err := us.userRepository.GetById(id)
	if err != nil {
		return GetByIdResponse{}, err
	}

	response := GetByIdResponseModelToResponse(user)

	return response, nil
}

func (us userService) CreateUser(name, email, password string, registerNumber int64, registerTypeID, userTypeID int) (string, error) {
	user, err := us.userRepository.CreateUser(name, email, password, registerNumber, registerTypeID, userTypeID)
	if err != nil {
		return "", err
	}

	fmt.Println(user)
	return "userCreated", nil
}
