package service

import (
	"github.com/aaguero96/technical_challenge_q2bank/repository"
	"github.com/aaguero96/technical_challenge_q2bank/service/register_type_service"
	"github.com/aaguero96/technical_challenge_q2bank/service/transaction_service"
	"github.com/aaguero96/technical_challenge_q2bank/service/user_service"
	"github.com/aaguero96/technical_challenge_q2bank/service/user_type_service"
	"github.com/aaguero96/technical_challenge_q2bank/service/wallet_service"
)

type Services struct {
	UserService         user_service.UserService
	UserTypeService     user_type_service.UserTypeService
	RegisterTypeService register_type_service.RegisterTypeService
	TransactionService  transaction_service.TransactionService
	WalletService       wallet_service.WalletService
}

func InstanceServices(repositories repository.Repositories) Services {
	userService := user_service.NewUserService(repositories.UserRepository)
	userTypeService := user_type_service.NewUserTypeService(repositories.UserTypeRepository)
	registerTypeService := register_type_service.NewRegisterTypeService(repositories.RegisterTypeRepository)
	transactionService := transaction_service.NewTransactionService(
		repositories.TransactionRepository,
		repositories.UserRepository,
		repositories.WalletRepository,
		repositories.UserTypeRepository,
	)
	walletService := wallet_service.NewWalletService(repositories.WalletRepository)

	return Services{
		UserService:         userService,
		UserTypeService:     userTypeService,
		RegisterTypeService: registerTypeService,
		TransactionService:  transactionService,
		WalletService:       walletService,
	}
}
