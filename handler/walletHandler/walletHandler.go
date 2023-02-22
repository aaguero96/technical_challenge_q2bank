package walletHandler

import (
	"errors"
	"net/http"
	"os"
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
// @Security 					BearerToken
// @Summary						Get all wallets
// @Description 			Get all wallets
// @Produce 					json
// @Tags 							wallet
// @Router						/v1/wallets [get]
// @Success						200 {object} []walletService.WalletResponse
// @Failure						500 {error} error
func (wh walletHandler) GetAll(ctx *gin.Context) {
	wallets, err := wh.walletService.GetAll()
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, wallets)
}

// GetById							godoc
// @Security 						BearerToken
// @Summary							Get wallet by id
// @Description 				Get wallet by id
// @Produce 						json
// @Tags 								wallet
// @Param   						id path int true "wallet id"
// @Router							/v1/wallets/{id} [get]
// @Success							200 {object} userTypeService.GetByIdResponse
// @Failure							400 {error} error
// @Failure							500 {error} error
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

// AddAmount						godoc
// @Security 						BasicAuth
// @Summary							Add amount in wallet
// @Description 				Add amount or descrease amount (in case of negative amount) in wallet (only for admin in moment)
// @Produce 						json
// @Tags 								wallet
// @Param   						id path int true "wallet id"
// @Param   						amount body AddAmountRequest true "Amount"
// @Router							/v1/wallets/{id} [patch]
// @Success							200 {object} models.WalletModel
// @Failure							400 {error} error
// @Failure							401 {error} error
// @Failure							500 {error} error
func (wh walletHandler) AddAmount(ctx *gin.Context) {
	// Verify if is admin
	username, password, ok := ctx.Request.BasicAuth()
	if !ok {
		httputil.NewError(ctx, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}
	if username != os.Getenv("ADMIN_USERNAME") || password != os.Getenv("ADMIN_PASSWORD") {
		httputil.NewError(ctx, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}

	paramID := ctx.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	var input AddAmountRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	wallet, err := wh.walletService.AddAmount(id, input.Amount)
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, wallet)
}
