package transactionHandler

import "github.com/gin-gonic/gin"

type TransactionHandler interface {
	GetAll(ctx *gin.Context)
}
