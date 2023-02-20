package userService

import (
	"github.com/aaguero96/technical_challenge_q2bank/repository/userRepository"
)

type userService struct {
	userRepository userRepository.UserRepository
}

func NewUserService(userRepository userRepository.UserRepository) userService {
	return userService{
		userRepository: userRepository,
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
