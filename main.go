package main

import (
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
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	// Start Repositories
	userRepository := userRepository.NewUserRepository(initializers.DB)
	userTypeRepository := userTypeRepository.NewUserTypeRepository(initializers.DB)
	registerTypeRepository := registerTypeRepository.NewRegisterTypeRepository(initializers.DB)
	transactionRepository := transactionRepository.NewTransactionRepository(initializers.DB)
	walletRepository := walletRepository.NewWalletRepository(initializers.DB)

	// Start Services
	userService := userService.NewUserService(userRepository)
	userTypeService := userTypeService.NewUserTypeService(userTypeRepository)
	registerTypeService := registerTypeService.NewRegisterTypeService(registerTypeRepository)
	transactionService := transactionService.NewTransactionService(transactionRepository)
	walletService := walletService.NewWalletService(walletRepository)

	// Start Handlers
	userHandler := userHandler.NewUserHandler(userService)
	userTypeHandler := userTypeHandler.NewUserTypeHandler(userTypeService)
	registerTypeHandler := registerTypeHandler.NewRegisterTypeHandler(registerTypeService)
	transactionHandler := transactionHandler.NewTransactionHandler(transactionService)
	walletHandler := walletHandler.NewWalletHandler(walletService)

	router := gin.Default()

	v1 := router.Group("/v1")
	routes.NewUserRoutes(v1, userHandler)
	routes.NewUserTypeRoutes(v1, userTypeHandler)
	routes.NewRegisterTypeRoutes(v1, registerTypeHandler)
	routes.NewTransactionRoutes(v1, transactionHandler)
	routes.NewWalletRoutes(v1, walletHandler)

	router.Run()
}
