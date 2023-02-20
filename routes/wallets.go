package routes

import (
	"github.com/aaguero96/technical_challenge_q2bank/handler/walletHandler"
	"github.com/gin-gonic/gin"
)

func NewWalletRoutes(rg *gin.RouterGroup, wh walletHandler.WalletHandler) {
	users := rg.Group("/wallets")

	users.GET("/:id", wh.GetById)
	users.GET("/", wh.GetAll)
}
