package user_type_service

type UserTypeService interface {
	GetAll() ([]UserTypeResponse, error)
	GetById(id int) (GetByIdResponse, error)
}
