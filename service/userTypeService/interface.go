package userTypeService

type UserTypeService interface {
	GetAll() ([]UserTypeResponse, error)
}
