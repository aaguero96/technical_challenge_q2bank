package routes

import (
	"github.com/aaguero96/technical_challenge_q2bank/handler/transactionHandler"
	"github.com/aaguero96/technical_challenge_q2bank/middleware"
	"github.com/gin-gonic/gin"
)

func NewTransactionRoutes(rg *gin.RouterGroup, th transactionHandler.TransactionHandler) {
	transactions := rg.Group("/transactions")

	transactions.Use(middleware.Authorization)
	transactions.GET("/", th.GetAll)
	transactions.POST("/", th.CreateTransaction)
}
