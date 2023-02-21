package routes

import (
	"github.com/aaguero96/technical_challenge_q2bank/handler/registerTypeHandler"
	"github.com/aaguero96/technical_challenge_q2bank/middleware"
	"github.com/gin-gonic/gin"
)

func NewRegisterTypeRoutes(rg *gin.RouterGroup, rth registerTypeHandler.RegisterTypeHandler) {
	registerTypes := rg.Group("/register_types")

	registerTypes.Use(middleware.Authorization)
	registerTypes.GET("/", rth.GetAll)
}
