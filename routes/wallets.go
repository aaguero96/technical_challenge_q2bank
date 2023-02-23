package routes

import (
	"github.com/aaguero96/technical_challenge_q2bank/handler/wallet_handler"
	"github.com/aaguero96/technical_challenge_q2bank/middleware"
	"github.com/gin-gonic/gin"
)

func NewWalletRoutes(rg *gin.RouterGroup, wh wallet_handler.WalletHandler) {
	wallets := rg.Group("/wallets")

	wallets.PATCH("/:id", wh.AddAmount)

	wallets.Use(middleware.Authorization)
	wallets.GET("/:id", wh.GetById)
	wallets.GET("/", wh.GetAll)
}
