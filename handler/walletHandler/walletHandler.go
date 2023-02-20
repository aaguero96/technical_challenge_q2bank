package walletHandler

import (
	"net/http"
	"strconv"

	"github.com/aaguero96/technical_challenge_q2bank/service/walletService"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
)

type walletHandler struct {
	walletService walletService.WalletService
}

func NewWalletHandler(ws walletService.WalletService) walletHandler {
	return walletHandler{
		walletService: ws,
	}
}

func (wh walletHandler) GetAll(ctx *gin.Context) {
	wallets, err := wh.walletService.GetAll()
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, wallets)
}

func (wh walletHandler) GetById(ctx *gin.Context) {
	paramID := ctx.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	user, err := wh.walletService.GetById(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}
