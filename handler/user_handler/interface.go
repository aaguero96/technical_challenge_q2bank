package user_handler

import "github.com/gin-gonic/gin"

type UserHandler interface {
	GetAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
}
