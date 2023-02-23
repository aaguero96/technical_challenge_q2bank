package user_type_service

import (
	"github.com/aaguero96/technical_challenge_q2bank/repository/user_type_repository"
)

type userTypeService struct {
	userTypeRepository user_type_repository.UserTypeRepository
}

func NewUserTypeService(utr user_type_repository.UserTypeRepository) userTypeService {
	return userTypeService{
		userTypeRepository: utr,
	}
}

func (uts userTypeService) GetAll() ([]UserTypeResponse, error) {
	userTypes, err := uts.userTypeRepository.GetAll()
	if err != nil {
		return nil, err
	}

	response := GetAllModelToResponse(userTypes)

	return response, nil
}

func (uts userTypeService) GetById(id int) (GetByIdResponse, error) {
	user, err := uts.userTypeRepository.GetById(id)
	if err != nil {
		return GetByIdResponse{}, err
	}

	response := GetByIdResponseModelToResponse(user)

	return response, nil
}
