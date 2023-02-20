package registerTypeHandler

import (
	"net/http"

	"github.com/aaguero96/technical_challenge_q2bank/service/registerTypeService"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
)

type registerTypeHandler struct {
	registerTypeService registerTypeService.RegisterTypeService
}

func NewRegisterTypeHandler(rts registerTypeService.RegisterTypeService) registerTypeHandler {
	return registerTypeHandler{
		registerTypeService: rts,
	}
}

func (rth registerTypeHandler) GetAll(ctx *gin.Context) {
	registerTypes, err := rth.registerTypeService.GetAll()
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, registerTypes)
}
