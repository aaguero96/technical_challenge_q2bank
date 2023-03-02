package utils

import (
	"github.com/aaguero96/technical_challenge_q2bank/initializers"
)

func ResetDatabase() {
	SetEnvs()
	initializers.ConnectDB()
	initializers.CreateDatabase()
}
