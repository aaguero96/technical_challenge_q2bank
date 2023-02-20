package registerTypeService

import "github.com/aaguero96/technical_challenge_q2bank/repository/registerTypeRepository"

type registerTypeService struct {
	registerTypeRepository registerTypeRepository.RegisterTypeRepository
}

func NewRegisterTypeService(rtr registerTypeRepository.RegisterTypeRepository) registerTypeService {
	return registerTypeService{
		registerTypeRepository: rtr,
	}
}

func (rts registerTypeService) GetAll() ([]RegiterTypeResponse, error) {
	registerTypes, err := rts.registerTypeRepository.GetAll()
	if err != nil {
		return nil, err
	}

	response := GetAllModelToResponse(registerTypes)

	return response, nil
}
