package initializers

import (
	"github.com/aaguero96/technical_challenge_q2bank/migrate/seeder"
	"github.com/aaguero96/technical_challenge_q2bank/models"
)

func CreateDatabase() {
	// Drop Tables
	DB.Migrator().DropTable(&models.RegisterTypeModel{})
	DB.Migrator().DropTable(&models.WalletModel{})
	DB.Migrator().DropTable(&models.UserTypeModel{})
	DB.Migrator().DropTable(&models.UserModel{})
	DB.Migrator().DropTable(&models.TransactionModel{})

	// Create tables
	DB.AutoMigrate(&models.RegisterTypeModel{})
	DB.AutoMigrate(&models.WalletModel{})
	DB.AutoMigrate(&models.UserTypeModel{})
	DB.AutoMigrate(&models.UserModel{})
	DB.AutoMigrate(&models.TransactionModel{})

	// Add values to tables
	DB.Model(&models.UserModel{}).Create(&seeder.UserSeeder)
	DB.Model(&models.RegisterTypeModel{}).Create(&seeder.RegisterTypeSeeder)
	DB.Model(&models.UserTypeModel{}).Create(&seeder.UserTypeSeeder)
	DB.Model(&models.WalletModel{}).Create(&seeder.WalletSeeder)
}
