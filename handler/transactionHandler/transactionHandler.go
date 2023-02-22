package transactionHandler

import (
	"net/http"
	"strconv"

	_ "github.com/aaguero96/technical_challenge_q2bank/docs"

	"github.com/aaguero96/technical_challenge_q2bank/service/transactionService"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
)

type transactionHandler struct {
	transactionService transactionService.TransactionService
}

func NewTransactionHandler(ts transactionService.TransactionService) transactionHandler {
	return transactionHandler{
		transactionService: ts,
	}
}

// GetAll							godoc
// @Summary						Get all transactions
// @Description 			Get all transactions
// @Produce 					json
// @Tags 							transaction
// @Router						/v1/transactions [get]
// @Success						200 {object} []transactionService.TransactionResponse
// @Success						500 {error} error
func (th transactionHandler) GetAll(ctx *gin.Context) {
	transactions, err := th.transactionService.GetAll()
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}

// CreateTransaction		godoc
// @Summary							Create transaction
// @Description 				Create transaction and send to redis queue to be aproved
// @Produce 						json
// @Tags 								transaction
// @Param   						transaction body CreateTransactionRequest true "Transaction data"
// @Router							/v1/transactions [post]
// @Success							201 {object} transactionService.CreateTransactionResponse
// @Success							400 {error} error
// @Success							500 {error} error
func (th transactionHandler) CreateTransaction(ctx *gin.Context) {
	payeerEmail := ctx.Request.Header.Get("email")

	var input CreateTransactionRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	transactions, err := th.transactionService.CreateTransaction(input.PayerID, input.PayeeID, input.Amount, payeerEmail)
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, transactions)
}

// CancelTransaction		godoc
// @Summary							Cancel transaction
// @Description 				Cancel transaction and send to redis queue to be aproved
// @Produce 						json
// @Tags 								transaction
// @Param   						id path int true "transaction id"
// @Router							/v1/transactions/{id} [delete]
// @Success							204
// @Success							400 {error} error
// @Success							500 {error} error
func (th transactionHandler) CancelTransaction(ctx *gin.Context) {
	paramID := ctx.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	payeerEmail := ctx.Request.Header.Get("email")

	err = th.transactionService.CancelTransaction(id, payeerEmail)
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	type Response struct {
		Message string `json:"message"`
	}

	ctx.JSON(http.StatusNoContent, Response{Message: "Cancel in progress"})
}

// GetById		godoc
// @Summary							Get transaction by id
// @Description 				Get transaction by id
// @Produce 						json
// @Tags 								transaction
// @Param   						id path int true "transaction id"
// @Router							/v1/transactions/{id} [get]
// @Success							200 {object} transactionService.GetByIdResponse
// @Success							400 {error} error
// @Success							500 {error} error
func (th transactionHandler) GetById(ctx *gin.Context) {
	paramID := ctx.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	transaction, err := th.transactionService.GetById(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, transaction)
}
