package handler

import (
	"net/http"

	"github.com/aaguero96/technical_challenge_q2bank/service/walletService"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
)

type walletHandler struct {
	walletService walletService.WalletService
}

func NewWalletHandler(walletService walletService.WalletService) walletHandler {
	return walletHandler{
		walletService: walletService,
	}
}

func (wh walletHandler) GetAll(ctx *gin.Context) {
	wallets, err := wh.walletService.GetAll()
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, wallets)
}
