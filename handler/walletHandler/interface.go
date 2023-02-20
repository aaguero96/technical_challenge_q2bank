package walletHandler

import "github.com/gin-gonic/gin"

type WalletHandler interface {
	GetAll(ctx *gin.Context)
}
