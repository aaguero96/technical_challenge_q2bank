package repository

import (
	"github.com/aaguero96/technical_challenge_q2bank/repository/register_type_repository"
	"github.com/aaguero96/technical_challenge_q2bank/repository/transaction_repository"
	"github.com/aaguero96/technical_challenge_q2bank/repository/user_repository"
	"github.com/aaguero96/technical_challenge_q2bank/repository/user_type_repository"
	"github.com/aaguero96/technical_challenge_q2bank/repository/wallet_repository"
	"gorm.io/gorm"
)

type Repositories struct {
	UserRepository         user_repository.UserRepository
	UserTypeRepository     user_type_repository.UserTypeRepository
	RegisterTypeRepository register_type_repository.RegisterTypeRepository
	TransactionRepository  transaction_repository.TransactionRepository
	WalletRepository       wallet_repository.WalletRepository
}

func InstanceRepositories(db *gorm.DB) Repositories {
	userRepository := user_repository.NewUserRepository(db)
	userTypeRepository := user_type_repository.NewUserTypeRepository(db)
	registerTypeRepository := register_type_repository.NewRegisterTypeRepository(db)
	transactionRepository := transaction_repository.NewTransactionRepository(db)
	walletRepository := wallet_repository.NewWalletRepository(db)

	return Repositories{
		UserRepository:         &userRepository,
		UserTypeRepository:     userTypeRepository,
		RegisterTypeRepository: registerTypeRepository,
		TransactionRepository:  transactionRepository,
		WalletRepository:       &walletRepository,
	}
}
