package transactionHandler

import "github.com/gin-gonic/gin"

type TransactionHandler interface {
	GetAll(ctx *gin.Context)
	CreateTransaction(ctx *gin.Context)
	CancelTransaction(ctx *gin.Context)
	GetById(ctx *gin.Context)
}
