package userService

type UserService interface {
	GetAll() ([]UserResponse, error)
}
