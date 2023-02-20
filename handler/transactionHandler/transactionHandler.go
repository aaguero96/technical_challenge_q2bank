package transactionHandler

import (
	"net/http"

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
	}

	ctx.JSON(http.StatusOK, transactions)
}
