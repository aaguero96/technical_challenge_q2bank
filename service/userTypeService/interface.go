package userTypeService

type UserTypeService interface {
	GetAll() ([]UserTypeResponse, error)
	GetById(id int) (GetByIdResponse, error)
}
