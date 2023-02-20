package userTypeService

import (
	"github.com/aaguero96/technical_challenge_q2bank/repository/userTypeRepository"
)

type userTypeService struct {
	userTypeRepository userTypeRepository.UserTypeRepository
}

func NewUserTypeService(utr userTypeRepository.UserTypeRepository) userTypeService {
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
