package userHandler

import (
	"net/http"

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
