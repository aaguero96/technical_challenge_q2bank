package main

import (
	"github.com/aaguero96/technical_challenge_q2bank/initializers"
	"github.com/aaguero96/technical_challenge_q2bank/migrate/seeder"
	"github.com/aaguero96/technical_challenge_q2bank/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	// Drop Tables
	initializers.DB.Migrator().DropTable(&models.RegisterTypeModel{})
	initializers.DB.Migrator().DropTable(&models.WalletModel{})
	initializers.DB.Migrator().DropTable(&models.UserTypeModel{})
	initializers.DB.Migrator().DropTable(&models.UserModel{})
	initializers.DB.Migrator().DropTable(&models.TransactionModel{})

	// Create tables
	initializers.DB.AutoMigrate(&models.RegisterTypeModel{})
	initializers.DB.AutoMigrate(&models.WalletModel{})
	initializers.DB.AutoMigrate(&models.UserTypeModel{})
	initializers.DB.AutoMigrate(&models.UserModel{})
	initializers.DB.AutoMigrate(&models.TransactionModel{})

	// Add values to tables
	initializers.DB.Model(&models.UserModel{}).Create(&seeder.UserSeeder)
	initializers.DB.Model(&models.RegisterTypeModel{}).Create(&seeder.RegisterTypeSeeder)
	initializers.DB.Model(&models.UserTypeModel{}).Create(&seeder.UserTypeSeeder)
	initializers.DB.Model(&models.WalletModel{}).Create(&seeder.WalletSeeder)
}
