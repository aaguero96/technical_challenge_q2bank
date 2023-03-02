package integration_test

import (
	"net/http"
	"testing"

	"github.com/aaguero96/technical_challenge_q2bank/integration_test/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewUser(t *testing.T) {
	assert := assert.New(t)

	// type Body for these requests
	type Body struct {
		Name           string `json:"name"`
		Email          string `json:"email"`
		Password       string `json:"password"`
		RegisterNumber int    `json:"register_number"`
		RegisterTypeID int    `json:"register_type_id"`
		UserTypeID     int    `json:"user_type_id"`
	}

	t.Run("Create user correctly but not accept cookies - OK CASE", func(t *testing.T) {
		// body
		body := Body{
			Name:           "name_0",
			Email:          "name0@test.com",
			Password:       "Def4!t*0",
			RegisterNumber: 12345678900,
			RegisterTypeID: 1,
			UserTypeID:     1,
		}

		// type reponse
		type responseType struct {
			Token    string `json:"token"`
			ExpingIn string `json:"expiring_in"`
		}

		// response
		status, data, cookies := utils.EndpointResponse[responseType](
			utils.EndpointRequest{Body: body}, "POST", utils.BASE_URL+"/v1/users?agree_cookie=false",
		)

		// assertions
		assert.Equal(http.StatusCreated, status)
		assert.Equal("30 minutes", data.ExpingIn)
		assert.NotEmpty(data.Token)
		assert.Equal(len(cookies), 0)

		// Reset database
		utils.ResetDatabase()
	})

	t.Run("Create user correctly and accept cookies - OK CASE", func(t *testing.T) {
		// body
		body := Body{
			Name:           "name_0",
			Email:          "name0@test.com",
			Password:       "Def4!t*0",
			RegisterNumber: 12345678900,
			RegisterTypeID: 1,
			UserTypeID:     1,
		}

		// type reponse
		type responseType struct {
			Token    string `json:"token"`
			ExpingIn string `json:"expiring_in"`
		}

		// response
		status, data, cookies := utils.EndpointResponse[responseType](
			utils.EndpointRequest{Body: body}, "POST", utils.BASE_URL+"/v1/users?agree_cookie=true",
		)

		// assertions
		assert.Equal(http.StatusCreated, status)
		assert.Equal("30 minutes", data.ExpingIn)
		assert.NotEmpty(data.Token)
		assert.Equal(cookies[0].Name, "token")
		assert.NotEmpty(cookies[0].Value)

		// Reset database
		utils.ResetDatabase()
	})
}
