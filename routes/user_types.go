package routes

import (
	"github.com/aaguero96/technical_challenge_q2bank/handler/user_type_handler"
	"github.com/aaguero96/technical_challenge_q2bank/middleware"
	"github.com/gin-gonic/gin"
)

func NewUserTypeRoutes(rg *gin.RouterGroup, uth user_type_handler.UserTypeHandler) {
	userTypes := rg.Group("/user-types")

	userTypes.Use(middleware.Authorization)
	userTypes.GET("/:id", uth.GetById)
	userTypes.GET("/", uth.GetAll)
}
