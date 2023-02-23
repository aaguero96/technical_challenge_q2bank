package register_type_handler

import "github.com/gin-gonic/gin"

type RegisterTypeHandler interface {
	GetAll(ctx *gin.Context)
}
