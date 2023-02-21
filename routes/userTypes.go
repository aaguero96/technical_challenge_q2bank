package routes

import (
	"github.com/aaguero96/technical_challenge_q2bank/handler/userTypeHandler"
	"github.com/gin-gonic/gin"
)

func NewUserTypeRoutes(rg *gin.RouterGroup, uth userTypeHandler.UserTypeHandler) {
	userTypes := rg.Group("/user_types")

	userTypes.GET("/:id", uth.GetById)
	userTypes.GET("/", uth.GetAll)
}
