package wallet_handler

import "github.com/gin-gonic/gin"

type WalletHandler interface {
	GetAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
	AddAmount(ctx *gin.Context)
}
