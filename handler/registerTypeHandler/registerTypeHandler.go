package registerTypeHandler

import (
	"net/http"

	_ "github.com/aaguero96/technical_challenge_q2bank/docs"

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

// GetAll							godoc
// @Security 					BearerToken
// @Summary						Get all regiter types
// @Description 			Get all register types
// @Produce 					json
// @Tags 							register type
// @Router						/v1/register_types [get]
// @Success						200 {object} []registerTypeService.RegisterTypeResponse
// @Failure						500 {error} error
func (rth registerTypeHandler) GetAll(ctx *gin.Context) {
	registerTypes, err := rth.registerTypeService.GetAll()
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, registerTypes)
}
