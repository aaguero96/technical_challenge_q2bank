package user_type_handler

import "github.com/gin-gonic/gin"

type UserTypeHandler interface {
	GetAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
}
