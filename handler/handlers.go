package handler

import (
	"github.com/aaguero96/technical_challenge_q2bank/handler/register_type_handler"
	"github.com/aaguero96/technical_challenge_q2bank/handler/transaction_handler"
	"github.com/aaguero96/technical_challenge_q2bank/handler/user_handler"
	"github.com/aaguero96/technical_challenge_q2bank/handler/user_type_handler"
	"github.com/aaguero96/technical_challenge_q2bank/handler/wallet_handler"
	"github.com/aaguero96/technical_challenge_q2bank/service"
)

type Handlers struct {
	UserHandler         user_handler.UserHandler
	UserTypeHandler     user_type_handler.UserTypeHandler
	RegisterTypeHandler register_type_handler.RegisterTypeHandler
	TransactionHandler  transaction_handler.TransactionHandler
	WalletHandler       wallet_handler.WalletHandler
}

func InstanceHandlers(services service.Services) Handlers {
	userHandler := user_handler.NewUserHandler(services.UserService)
	userTypeHandler := user_type_handler.NewUserTypeHandler(services.UserTypeService)
	registerTypeHandler := register_type_handler.NewRegisterTypeHandler(services.RegisterTypeService)
	transactionHandler := transaction_handler.NewTransactionHandler(services.TransactionService)
	walletHandler := wallet_handler.NewWalletHandler(services.WalletService)

	return Handlers{
		UserHandler:         &userHandler,
		UserTypeHandler:     userTypeHandler,
		RegisterTypeHandler: registerTypeHandler,
		TransactionHandler:  transactionHandler,
		WalletHandler:       &walletHandler,
	}
}
