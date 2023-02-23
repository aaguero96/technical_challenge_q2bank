package user_type_handler

import (
	"net/http"
	"strconv"

	"github.com/aaguero96/technical_challenge_q2bank/service/user_type_service"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
)

type userTypeHandler struct {
	userTypeService user_type_service.UserTypeService
}

func NewUserTypeHandler(uts user_type_service.UserTypeService) userTypeHandler {
	return userTypeHandler{
		userTypeService: uts,
	}
}

// GetAll							godoc
// @Security 					BearerToken
// @Summary						Get all user types
// @Description 			Get all user types
// @Produce 					json
// @Tags 							user type
// @Router						/v1/user-types [get]
// @Success						200 {object} []user_type_service.UserTypeResponse
// @Failure						500 {error} error
func (uth userTypeHandler) GetAll(ctx *gin.Context) {
	userTypes, err := uth.userTypeService.GetAll()
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, userTypes)
}

// GetById							godoc
// @Security 						BearerToken
// @Summary							Get user type by id
// @Description 				Get user type by id
// @Produce 						json
// @Tags 								user type
// @Param   						id path int true "user type id"
// @Router							/v1/user-types/{id} [get]
// @Success							200 {object} user_type_service.GetByIdResponse
// @Failure							400 {error} error
// @Failure							500 {error} error
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
