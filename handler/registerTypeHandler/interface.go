package registerTypeHandler

import "github.com/gin-gonic/gin"

type RegisterTypeHandler interface {
	GetAll(ctx *gin.Context)
}
