package userTypeService

type UserService interface {
	GetAll() ([]UserTypeResponse, error)
}
