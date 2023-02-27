package integration_test

import (
	"net/http"
	"testing"

	"github.com/aaguero96/technical_challenge_q2bank/integration_test/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetAllWallets(t *testing.T) {
	assert := assert.New(t)

	t.Run("Create user correctly but not accept cookies - OK CASE", func(t *testing.T) {
		// create user
		token := utils.CreateUser(
			"name_1",
			"name1@test.com",
			"Def4!t*01",
			12345678901,
			1,
			1,
		)

		// type reponse
		type responseType struct {
			ID     int     `json:"id"`
			Amount float64 `json:"amount"`
		}

		// response
		status, data, _ := utils.EndpointResponse[[]responseType](
			utils.EndpointRequest{
				Body: nil,
				Header: utils.EndpointHeader{
					BearerToken: "Bearer " + token,
				},
			},
			"GET",
			"http://0.0.0.0:3000/v1/wallets",
		)

		// assertions
		assert.Equal(http.StatusOK, status)
		assert.Equal([]responseType{
			{ID: 1, Amount: 10000},
			{ID: 2, Amount: 20000},
			{ID: 3, Amount: 0},
		}, data)

		// Reset database
		utils.ResetDatabase()
	})
}
