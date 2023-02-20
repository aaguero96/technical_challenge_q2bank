package userTypeHandler

import "github.com/gin-gonic/gin"

type UserTypeHandler interface {
	GetAll(ctx *gin.Context)
}
