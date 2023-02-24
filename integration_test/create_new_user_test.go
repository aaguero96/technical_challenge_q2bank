package integration_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

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
		body, err := json.Marshal(Body{
			Name:           "name_0",
			Email:          "name0@test.com",
			Password:       "Def4!t*0",
			RegisterNumber: 12345678900,
			RegisterTypeID: 1,
			UserTypeID:     1,
		})
		if err != nil {
			fmt.Println("error when parsisng json")
		}

		// http request
		request, err := http.NewRequest("POST", "http://0.0.0.0:3000/v1/users/", bytes.NewBuffer(body))
		if err != nil {
			fmt.Println("error when request is created")
		}

		// http response
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			fmt.Println("error when response is created")
		}
		defer response.Body.Close()

		// verify status code
		assert.Equal(http.StatusCreated, response.StatusCode, "expected status %v, instead got %v", http.StatusCreated, response.StatusCode)

		// // verify Response
		// type reponse
		type responseType struct {
			Token    string `json:"token"`
			ExpingIn string `json:"expiring_in"`
		}

		// read response data
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("error to read response")
		}

		// verify if data has correct type
		var data responseType
		if err = json.Unmarshal([]byte(string(responseData)), &data); err != nil {
			fmt.Println("data dont has correct types")
		}

		// verify if token not equal to ""
		assert.NotEmpty(data.Token)

		// verify if expiredIn is correct
		assert.Equal("30 minutes", data.ExpingIn)
	})
}
