package routes

import (
	"github.com/aaguero96/technical_challenge_q2bank/handler/user_handler"
	"github.com/gin-gonic/gin"
)

func NewLoginRoutes(rg *gin.RouterGroup, uh user_handler.UserHandler) {
	login := rg.Group("/login")

	login.POST("/", uh.LoginUser)
}
