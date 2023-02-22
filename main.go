package main

import (
	_ "github.com/aaguero96/technical_challenge_q2bank/docs"
	"github.com/aaguero96/technical_challenge_q2bank/handler/registerTypeHandler"
	"github.com/aaguero96/technical_challenge_q2bank/handler/transactionHandler"
	"github.com/aaguero96/technical_challenge_q2bank/handler/userHandler"
	"github.com/aaguero96/technical_challenge_q2bank/handler/userTypeHandler"
	"github.com/aaguero96/technical_challenge_q2bank/handler/walletHandler"
	"github.com/aaguero96/technical_challenge_q2bank/initializers"
	"github.com/aaguero96/technical_challenge_q2bank/repository/registerTypeRepository"
	"github.com/aaguero96/technical_challenge_q2bank/repository/transactionRepository"
	"github.com/aaguero96/technical_challenge_q2bank/repository/userRepository"
	"github.com/aaguero96/technical_challenge_q2bank/repository/userTypeRepository"
	"github.com/aaguero96/technical_challenge_q2bank/repository/walletRepository"
	"github.com/aaguero96/technical_challenge_q2bank/routes"
	"github.com/aaguero96/technical_challenge_q2bank/service/registerTypeService"
	"github.com/aaguero96/technical_challenge_q2bank/service/transactionService"
	"github.com/aaguero96/technical_challenge_q2bank/service/userService"
	"github.com/aaguero96/technical_challenge_q2bank/service/userTypeService"
	"github.com/aaguero96/technical_challenge_q2bank/service/walletService"
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
func main() {
	// Start Repositories
	userRepository := userRepository.NewUserRepository(initializers.DB)
	userTypeRepository := userTypeRepository.NewUserTypeRepository(initializers.DB)
	registerTypeRepository := registerTypeRepository.NewRegisterTypeRepository(initializers.DB)
	transactionRepository := transactionRepository.NewTransactionRepository(initializers.DB)
	walletRepository := walletRepository.NewWalletRepository(initializers.DB)

	// Start Services
	userService := userService.NewUserService(&userRepository)
	userTypeService := userTypeService.NewUserTypeService(userTypeRepository)
	registerTypeService := registerTypeService.NewRegisterTypeService(registerTypeRepository)
	transactionService := transactionService.NewTransactionService(transactionRepository, &userRepository, walletRepository, userTypeRepository)
	walletService := walletService.NewWalletService(walletRepository)

	// Start Handlers
	userHandler := userHandler.NewUserHandler(userService)
	userTypeHandler := userTypeHandler.NewUserTypeHandler(userTypeService)
	registerTypeHandler := registerTypeHandler.NewRegisterTypeHandler(registerTypeService)
	transactionHandler := transactionHandler.NewTransactionHandler(transactionService)
	walletHandler := walletHandler.NewWalletHandler(walletService)

	router := gin.Default()

	// Add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/v1")

	routes.NewLoginRoutes(v1, userHandler)
	routes.NewUserRoutes(v1, userHandler)
	routes.NewUserTypeRoutes(v1, userTypeHandler)
	routes.NewRegisterTypeRoutes(v1, registerTypeHandler)
	routes.NewTransactionRoutes(v1, transactionHandler)
	routes.NewWalletRoutes(v1, walletHandler)

	router.Run()
}
