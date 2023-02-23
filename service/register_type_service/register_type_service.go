package register_type_service

import "github.com/aaguero96/technical_challenge_q2bank/repository/register_type_repository"

type registerTypeService struct {
	registerTypeRepository register_type_repository.RegisterTypeRepository
}

func NewRegisterTypeService(rtr register_type_repository.RegisterTypeRepository) registerTypeService {
	return registerTypeService{
		registerTypeRepository: rtr,
	}
}

func (rts registerTypeService) GetAll() ([]RegisterTypeResponse, error) {
	registerTypes, err := rts.registerTypeRepository.GetAll()
	if err != nil {
		return nil, err
	}

	response := GetAllModelToResponse(registerTypes)

	return response, nil
}
