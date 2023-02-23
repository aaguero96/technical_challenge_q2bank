package register_type_service

type RegisterTypeService interface {
	GetAll() ([]RegisterTypeResponse, error)
}
