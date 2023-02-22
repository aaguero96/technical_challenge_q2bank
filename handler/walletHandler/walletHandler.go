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

// GetAll							godoc
// @Summary						Get all wallets
// @Description 			Get all wallets
// @Produce 					json
// @Tags 							wallet
// @Router						/v1/wallets [get]
// @Success						200 {object} []walletService.WalletResponse
// @Success						500 {error} error
func (wh walletHandler) GetAll(ctx *gin.Context) {
	wallets, err := wh.walletService.GetAll()
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, wallets)
}

// GetById							godoc
// @Summary							Get wallet by id
// @Description 				Get wallet by id
// @Produce 						json
// @Tags 								wallet
// @Param   						id path int true "wallet id"
// @Router							/v1/wallets/{id} [get]
// @Success							200 {object} userTypeService.GetByIdResponse
// @Success							400 {error} error
// @Success							500 {error} error
func (wh walletHandler) GetById(ctx *gin.Context) {
	paramID := ctx.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	wallet, err := wh.walletService.GetById(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, wallet)
}
