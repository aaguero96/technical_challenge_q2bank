package main

import (
	"github.com/aaguero96/technical_challenge_q2bank/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	initializers.CreateDatabase()
}
