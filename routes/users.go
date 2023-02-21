package routes

import (
	"github.com/aaguero96/technical_challenge_q2bank/handler/userHandler"
	"github.com/gin-gonic/gin"
)

func NewUserRoutes(rg *gin.RouterGroup, uh userHandler.UserHandler) {
	users := rg.Group("/users")

	users.GET("/:id", uh.GetById)
	users.GET("/", uh.GetAll)
	users.POST("/", uh.CreateUser)
}
