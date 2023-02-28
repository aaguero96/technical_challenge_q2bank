package integration_test

import (
	"net/http"
	"testing"

	"github.com/aaguero96/technical_challenge_q2bank/integration_test/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetAllWallets(t *testing.T) {
	assert := assert.New(t)

	t.Run("Get all wallets correctly - OK CASE", func(t *testing.T) {
		// login user to get token
		token := utils.LoginUser("email1@testmail.com", "Def4!t*1")

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
			utils.BASE_URL+"/v1/wallets",
		)

		// assertions
		assert.Equal(http.StatusOK, status)
		assert.Equal([]responseType{
			{ID: 1, Amount: 100000},
			{ID: 2, Amount: 20000},
		}, data)
	})
}
