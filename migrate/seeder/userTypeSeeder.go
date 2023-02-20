package seeder

import "github.com/aaguero96/technical_challenge_q2bank/models"

var UserTypeSeeder = []models.UserTypeModel{
	{
		UserTypeID: 1,
		UserType:   "common",
	},
	{
		UserTypeID: 2,
		UserType:   "storekeeper",
	},
}
