package userHandler

type CreateUserRequest struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	RegisterNumber int64  `json:"register_number"`
	RegisterTypeID int    `json:"register_type_id"`
	UserTypeID     int    `json:"user_type_id"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
