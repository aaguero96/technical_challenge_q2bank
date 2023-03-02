package integration_test

import (
	"net/http"
	"testing"

	"github.com/aaguero96/technical_challenge_q2bank/integration_test/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	assert := assert.New(t)

	// type reponse
	type responseType struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		WalletID int    `json:"wallet_id"`
	}

	t.Run("Get all wallets correctly - OK CASE", func(t *testing.T) {
		// login user to get token
		token := utils.LoginUser("email1@testmail.com", "Def4!t*1")

		// response
		status, data, _ := utils.EndpointResponse[[]responseType](
			utils.EndpointRequest{
				Body: nil,
				Header: utils.EndpointHeader{
					BearerToken: "Bearer " + token,
				},
			},
			"GET",
			utils.BASE_URL+"/v1/users",
		)

		// assertions
		assert.Equal(http.StatusOK, status)
		assert.Equal([]responseType{
			{Name: "name_1_common", Email: "email1@testmail.com", WalletID: 1},
			{Name: "name_2_storekeeper", Email: "email2@testmail.com", WalletID: 2},
		}, data)
	})
}

func TestGetUserByID(t *testing.T) {
	assert := assert.New(t)

	// type reponse
	type responseType struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		WalletID int    `json:"wallet_id"`
	}

	t.Run("Get wallet by id correctly - OK CASE", func(t *testing.T) {
		// login user to get token
		token := utils.LoginUser("email1@testmail.com", "Def4!t*1")

		// response
		status, data, _ := utils.EndpointResponse[responseType](
			utils.EndpointRequest{
				Body: nil,
				Header: utils.EndpointHeader{
					BearerToken: "Bearer " + token,
				},
			},
			"GET",
			utils.BASE_URL+"/v1/users/1",
		)

		// assertions
		assert.Equal(http.StatusOK, status)
		assert.Equal(responseType{Name: "name_1_common", Email: "email1@testmail.com", WalletID: 1}, data)
	})
}
