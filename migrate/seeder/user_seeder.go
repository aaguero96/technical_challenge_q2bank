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
		Password:       "$2a$10$16yn5uLLPNhTcalc4XWDUuEgq.uiKywVEMa9SuM7KBN6iVGzMIYui", // without hash: Def4!t*1
	},
	{
		Name:           "name_2_storekeeper",
		RegisterNumber: 22222222222222,
		RegisterTypeID: 2,
		Email:          "email2@testmail.com",
		WalletID:       2,
		UserTypeID:     2,
		Password:       "$2a$10$WGJoxCGpj.rAm3CCU3z8pOXdgm4VCZhsOkrGIj7f2O2/DLwYupXKy", // without hash: Def4!t*2
	},
}
