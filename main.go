package main

import (
	_ "github.com/aaguero96/technical_challenge_q2bank/docs"
	"github.com/aaguero96/technical_challenge_q2bank/handler"
	"github.com/aaguero96/technical_challenge_q2bank/initializers"
	"github.com/aaguero96/technical_challenge_q2bank/repository"
	"github.com/aaguero96/technical_challenge_q2bank/routes"
	"github.com/aaguero96/technical_challenge_q2bank/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.ConnectRedisClient()
}

// @title Technical Challenge Q2bank
// @version 1.0
// @description A transaction service API in Go using Gin framework, Redis for queue works and Postgres as relational bank
// @host localhost:3000
// @BasePath /
// @securityDefinitions.apikey BearerToken
// @in header
// @name Authorization
// @securityDefinitions.basic BasicAuth
// @in header
// @name Admin Authorization
func main() {
	repositories := repository.InstanceRepositories(initializers.DB)
	service := service.InstanceServices(repositories)
	handlers := handler.InstanceHandlers(service)

	router := gin.Default()

	// Add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/v1")

	routes.NewLoginRoutes(v1, handlers.UserHandler)
	routes.NewUserRoutes(v1, handlers.UserHandler)
	routes.NewUserTypeRoutes(v1, handlers.UserTypeHandler)
	routes.NewRegisterTypeRoutes(v1, handlers.RegisterTypeHandler)
	routes.NewTransactionRoutes(v1, handlers.TransactionHandler)
	routes.NewWalletRoutes(v1, handlers.WalletHandler)

	router.Run()
}
