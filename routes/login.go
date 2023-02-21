package routes

import (
	"github.com/aaguero96/technical_challenge_q2bank/handler/userHandler"
	"github.com/gin-gonic/gin"
)

func NewLoginRoutes(rg *gin.RouterGroup, uh userHandler.UserHandler) {
	login := rg.Group("/login")

	login.POST("/", uh.LoginUser)
}
