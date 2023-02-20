package userService

type UserService interface {
	GetAll() ([]UserResponse, error)
	GetById(id int) (GetByIdResponse, error)
}
