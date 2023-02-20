package userService

import (
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
