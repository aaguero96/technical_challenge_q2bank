package transactionHandler

import (
	"net/http"
	"strconv"

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

func (th transactionHandler) GetAll(ctx *gin.Context) {
	transactions, err := th.transactionService.GetAll()
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}

func (th transactionHandler) CreateTransaction(ctx *gin.Context) {
	type request struct {
		PayerID int     `json:"payer_id"`
		PayeeID int     `json:"payee_id"`
		Amount  float64 `json:"amount"`
	}

	payeerEmail := ctx.Request.Header.Get("email")

	var input request
	if err := ctx.ShouldBindJSON(&input); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	transactions, err := th.transactionService.CreateTransaction(input.PayerID, input.PayeeID, input.Amount, payeerEmail)
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}

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

	ctx.JSON(http.StatusOK, Response{Message: "Cancel in progress"})
}
