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

func (uh userHandler) GetAll(ctx *gin.Context) {
	users, err := uh.userService.GetAll()
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, users)
}

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
