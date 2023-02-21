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

func (uth userTypeHandler) GetAll(ctx *gin.Context) {
	userTypes, err := uth.userTypeService.GetAll()
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, userTypes)
}

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
