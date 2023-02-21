package userService

type UserService interface {
	GetAll() ([]UserResponse, error)
	GetById(id int) (GetByIdResponse, error)
	CreateUser(name, email, password string, registerNumber int64, registerTypeID, userTypeID int) (CreateUserResponse, error)
	LoginUser(email, password string) (LoginUserResponse, error)
}
