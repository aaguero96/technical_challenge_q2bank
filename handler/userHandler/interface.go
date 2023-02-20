package userHandler

import "github.com/gin-gonic/gin"

type UserHandler interface {
	GetAll(ctx *gin.Context)
}
