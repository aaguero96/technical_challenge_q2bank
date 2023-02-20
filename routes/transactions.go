package routes

import (
	"github.com/aaguero96/technical_challenge_q2bank/handler/transactionHandler"
	"github.com/gin-gonic/gin"
)

func NewTransactionRoutes(rg *gin.RouterGroup, th transactionHandler.TransactionHandler) {
	users := rg.Group("/transactions")

	users.GET("/", th.GetAll)
	users.POST("/", th.CreateTransaction)
}
