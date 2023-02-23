package main

import (
	_ "github.com/aaguero96/technical_challenge_q2bank/docs"
	"github.com/aaguero96/technical_challenge_q2bank/handler/register_type_handler"
	"github.com/aaguero96/technical_challenge_q2bank/handler/transaction_handler"
	"github.com/aaguero96/technical_challenge_q2bank/handler/user_handler"
	"github.com/aaguero96/technical_challenge_q2bank/handler/user_type_handler"
	"github.com/aaguero96/technical_challenge_q2bank/handler/wallet_handler"
	"github.com/aaguero96/technical_challenge_q2bank/initializers"
	"github.com/aaguero96/technical_challenge_q2bank/repository/register_type_repository"
	"github.com/aaguero96/technical_challenge_q2bank/repository/transaction_repository"
	"github.com/aaguero96/technical_challenge_q2bank/repository/user_repository"
	"github.com/aaguero96/technical_challenge_q2bank/repository/user_type_repository"
	"github.com/aaguero96/technical_challenge_q2bank/repository/wallet_repository"
	"github.com/aaguero96/technical_challenge_q2bank/routes"
	"github.com/aaguero96/technical_challenge_q2bank/service/register_type_service"
	"github.com/aaguero96/technical_challenge_q2bank/service/transaction_service"
	"github.com/aaguero96/technical_challenge_q2bank/service/user_service"
	"github.com/aaguero96/technical_challenge_q2bank/service/user_type_service"
	"github.com/aaguero96/technical_challenge_q2bank/service/wallet_service"
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
	// Start Repositories
	userRepository := user_repository.NewUserRepository(initializers.DB)
	userTypeRepository := user_type_repository.NewUserTypeRepository(initializers.DB)
	registerTypeRepository := register_type_repository.NewRegisterTypeRepository(initializers.DB)
	transactionRepository := transaction_repository.NewTransactionRepository(initializers.DB)
	walletRepository := wallet_repository.NewWalletRepository(initializers.DB)

	// Start Services
	userService := user_service.NewUserService(&userRepository)
	userTypeService := user_type_service.NewUserTypeService(userTypeRepository)
	registerTypeService := register_type_service.NewRegisterTypeService(registerTypeRepository)
	transactionService := transaction_service.NewTransactionService(transactionRepository, &userRepository, &walletRepository, userTypeRepository)
	walletService := wallet_service.NewWalletService(&walletRepository)

	// Start Handlers
	userHandler := user_handler.NewUserHandler(userService)
	userTypeHandler := user_type_handler.NewUserTypeHandler(userTypeService)
	registerTypeHandler := register_type_handler.NewRegisterTypeHandler(registerTypeService)
	transactionHandler := transaction_handler.NewTransactionHandler(transactionService)
	walletHandler := wallet_handler.NewWalletHandler(walletService)

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
