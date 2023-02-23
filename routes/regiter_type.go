package routes

import (
	"github.com/aaguero96/technical_challenge_q2bank/handler/register_type_handler"
	"github.com/aaguero96/technical_challenge_q2bank/middleware"
	"github.com/gin-gonic/gin"
)

func NewRegisterTypeRoutes(rg *gin.RouterGroup, rth register_type_handler.RegisterTypeHandler) {
	registerTypes := rg.Group("/register_types")

	registerTypes.Use(middleware.Authorization)
	registerTypes.GET("/", rth.GetAll)
}
