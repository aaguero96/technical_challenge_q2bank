package userTypeHandler

import (
	"net/http"
	"strconv"

	"github.com/aaguero96/technical_challenge_q2bank/service/userTypeService"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
)

type userTypeHandler struct {
	userTypeService userTypeService.UserTypeService
}

func NewUserTypeHandler(uts userTypeService.UserTypeService) userTypeHandler {
	return userTypeHandler{
		userTypeService: uts,
	}
}

// GetAll							godoc
// @Summary						Get all user types
// @Description 			Get all user types
// @Produce 					json
// @Tags 							user type
// @Router						/v1/user_types [get]
// @Success						200 {object} []userTypeService.UserTypeResponse
// @Success						500 {error} error
func (uth userTypeHandler) GetAll(ctx *gin.Context) {
	userTypes, err := uth.userTypeService.GetAll()
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, userTypes)
}

// GetById							godoc
// @Summary							Get user type by id
// @Description 				Get user type by id
// @Produce 						json
// @Tags 								user type
// @Param   						id path int true "user type id"
// @Router							/v1/user_types/{id} [get]
// @Success							200 {object} userTypeService.GetByIdResponse
// @Success							400 {error} error
// @Success							500 {error} error
func (uth userTypeHandler) GetById(ctx *gin.Context) {
	paramID := ctx.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	userType, err := uth.userTypeService.GetById(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, userType)
}
