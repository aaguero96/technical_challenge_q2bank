package seeder

import (
	"github.com/aaguero96/technical_challenge_q2bank/models"
)

var UserSeeder = []models.UserModel{
	{
		Name:           "name_1_common",
		RegisterNumber: 11111111111,
		RegisterTypeID: 1,
		Email:          "email1@testmail.com",
		WalletID:       1,
		UserTypeID:     1,
	},
	{
		Name:           "name_2_storekeeper",
		RegisterNumber: 22222222222222,
		RegisterTypeID: 2,
		Email:          "email1@testmail.com",
		WalletID:       2,
		UserTypeID:     2,
	},
}
