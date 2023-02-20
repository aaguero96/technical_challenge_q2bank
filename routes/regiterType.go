package routes

import (
	"github.com/aaguero96/technical_challenge_q2bank/handler/registerTypeHandler"
	"github.com/gin-gonic/gin"
)

func NewRegisterTypeRoutes(rg *gin.RouterGroup, rth registerTypeHandler.RegisterTypeHandler) {
	users := rg.Group("/register_types")

	users.GET("/", rth.GetAll)
}
