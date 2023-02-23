package user_service

import (
	"errors"

	"github.com/aaguero96/technical_challenge_q2bank/repository/user_repository"
	"github.com/aaguero96/technical_challenge_q2bank/utils"
)

type userService struct {
	userRepository user_repository.UserRepository
}

func NewUserService(ur user_repository.UserRepository) userService {
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

func (us userService) CreateUser(name, email, password string, registerNumber int64, registerTypeID, userTypeID int) (CreateUserResponse, error) {
	// Validate Password
	if errorMessage, ok := utils.ValidatePasswordFormat(password); !ok {
		return CreateUserResponse{}, errors.New(errorMessage)
	}

	// Validate Email
	if !utils.ValidateEmail(email) {
		return CreateUserResponse{}, errors.New("email has not the correct format (eg: email@email.com)")
	}

	user, err := us.userRepository.CreateUser(name, email, password, registerNumber, registerTypeID, userTypeID)
	if err != nil {
		return CreateUserResponse{}, err
	}

	// Create JWT
	token, err := utils.CreateJWT(user.Email)
	if err != nil {
		return CreateUserResponse{}, err
	}

	return CreateUserResponse{
		Token:      token,
		ExpiringIn: "30 minutes",
	}, nil
}

func (us userService) LoginUser(email, password string) (LoginUserResponse, error) {
	user, err := us.userRepository.LoginUser(email, password)
	if err != nil {
		return LoginUserResponse{}, err
	}

	// Create JWT
	token, err := utils.CreateJWT(user.Email)
	if err != nil {
		return LoginUserResponse{}, err
	}

	return LoginUserResponse{
		Token:      token,
		ExpiringIn: "30 minutes",
	}, nil
}
