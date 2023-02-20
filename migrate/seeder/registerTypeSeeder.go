package seeder

import (
	"github.com/aaguero96/technical_challenge_q2bank/models"
)

var RegisterTypeSeeder = []models.RegisterTypeModel{
	{
		RegisterTypeID: 1,
		RegisterType:   "CPF",
	},
	{
		RegisterTypeID: 2,
		RegisterType:   "CNPJ2",
	},
}
