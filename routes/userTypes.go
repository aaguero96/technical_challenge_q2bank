package routes

import (
	"github.com/aaguero96/technical_challenge_q2bank/handler/userTypeHandler"
	"github.com/gin-gonic/gin"
)

func NewUserTypeRoutes(rg *gin.RouterGroup, uth userTypeHandler.UserTypeHandler) {
	users := rg.Group("/user_types")

	users.GET("/:id", uth.GetById)
	users.GET("/", uth.GetAll)
}
