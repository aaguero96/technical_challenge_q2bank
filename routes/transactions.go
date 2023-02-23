package routes

import (
	"github.com/aaguero96/technical_challenge_q2bank/handler/transaction_handler"
	"github.com/aaguero96/technical_challenge_q2bank/middleware"
	"github.com/gin-gonic/gin"
)

func NewTransactionRoutes(rg *gin.RouterGroup, th transaction_handler.TransactionHandler) {
	transactions := rg.Group("/transactions")

	transactions.Use(middleware.Authorization)
	transactions.GET("/:id", th.GetById)
	transactions.GET("/", th.GetAll)
	transactions.POST("/", th.CreateTransaction)
	transactions.DELETE("/:id", th.CancelTransaction)
}
