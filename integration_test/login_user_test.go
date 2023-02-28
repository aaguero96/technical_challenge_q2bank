package integration_test

import (
	"net/http"
	"testing"

	"github.com/aaguero96/technical_challenge_q2bank/integration_test/utils"
	"github.com/stretchr/testify/assert"
)

func TestLoginUser(t *testing.T) {
	assert := assert.New(t)

	// type Body for these requests
	type Body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	t.Run("Login user correctly but not accept cookies - OK CASE", func(t *testing.T) {
		// body
		body := Body{
			Email:    "email1@testmail.com",
			Password: "Def4!t*1",
		}

		// type reponse
		type responseType struct {
			Token    string `json:"token"`
			ExpingIn string `json:"expiring_in"`
		}

		// response
		status, data, cookies := utils.EndpointResponse[responseType](
			utils.EndpointRequest{Body: body}, "POST", "http://0.0.0.0:3000/v1/login?agree_cookie=false",
		)

		// assertions
		assert.Equal(http.StatusCreated, status)
		assert.Equal("30 minutes", data.ExpingIn)
		assert.NotEmpty(data.Token)
		assert.Equal(len(cookies), 0)
	})
}
