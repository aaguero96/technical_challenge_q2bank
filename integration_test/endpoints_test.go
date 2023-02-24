package integration_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndpoints(t *testing.T) {
	assert := assert.New(t)

	// Create http client
	client := &http.Client{}

	t.Run("Endpoint GET /users exist", func(t *testing.T) {
		// http request
		request, err := http.NewRequest("GET", "http://0.0.0.0:3000/v1/users/", nil)
		if err != nil {
			fmt.Println("error when request is created")
		}

		// http response
		response, err := client.Do(request)
		if err != nil {
			fmt.Println("error when response is created")
		}
		defer response.Body.Close()

		// Verify status code
		assert.NotEqual(http.StatusNotFound, response.StatusCode)

		// verify Response
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("error to read response")
		}
		var data string
		if err = json.Unmarshal([]byte(string(responseData)), &data); err != nil {
			fmt.Println("data dont has correct types")
		}
		assert.NotEqual("404 page not found", data)
	})
}
