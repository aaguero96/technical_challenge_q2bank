package userHandler

import (
	"net/http"
	"strconv"

	"github.com/aaguero96/technical_challenge_q2bank/service/userService"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
)

type userHandler struct {
	userService userService.UserService
}

func NewUserHandler(us userService.UserService) userHandler {
	return userHandler{
		userService: us,
	}
}

// GetAll							godoc
// @Summary						Get all users
// @Description 			Get all users
// @Produce 					json
// @Tags 							user
// @Router						/v1/users [get]
// @Success						200 {object} []userService.UserResponse
// @Success						500 {error} error
func (uh userHandler) GetAll(ctx *gin.Context) {
	users, err := uh.userService.GetAll()
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// GetById							godoc
// @Summary							Get users by id
// @Description 				Get users by id
// @Produce 						json
// @Tags 								user
// @Param   						id path int true "user id"
// @Router							/v1/users/{id} [get]
// @Success							200 {object} userService.GetByIdResponse
// @Success							400 {error} error
// @Success							500 {error} error
func (uh userHandler) GetById(ctx *gin.Context) {
	paramID := ctx.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	user, err := uh.userService.GetById(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// CreateUser						godoc
// @Summary							Create user
// @Description 				Create user
// @Produce 						json
// @Tags 								user
// @Param   						user body CreateUserRequest true "User data"
// @Router							/v1/users [post]
// @Success							200 {object} userService.CreateUserResponse
// @Success							400 {error} error
// @Success							500 {error} error
func (uh userHandler) CreateUser(ctx *gin.Context) {
	var input CreateUserRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	response, err := uh.userService.CreateUser(input.Name, input.Email, input.Password, input.RegisterNumber, input.RegisterTypeID, input.UserTypeID)
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.SetCookie("token", response.Token, 3600, "/", "localhost", false, true)
	ctx.JSON(http.StatusCreated, response)
}

// CreateUser						godoc
// @Summary							Create user
// @Description 				Create user
// @Produce 						json
// @Tags 								login
// @Param   						user body LoginRequest true "User credencial"
// @Router							/v1/login [post]
// @Success							200 {object} userService.LoginUserResponse
// @Success							400 {error} error
// @Success							500 {error} error
func (uh userHandler) LoginUser(ctx *gin.Context) {
	var input LoginRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	response, err := uh.userService.LoginUser(input.Email, input.Password)
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.SetCookie("token", response.Token, 3600, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, response)
}
