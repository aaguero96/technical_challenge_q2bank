package integration_test

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCreateNewUser(t *testing.T) {
	t.Run("Create user correctly but not accept cookies - OK CASE", func(t *testing.T) {
		request, err := http.NewRequest("POST", "localhost:3000/v1/users", nil)
		if err != nil {
			fmt.Println("error when request is created")
		}
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			fmt.Println("error when response is created")
		}
		defer response.Body.Close()

		fmt.Println("response status", response.Status)
		fmt.Println("response status code", response.StatusCode)
		fmt.Println("response body", response.Body)
		fmt.Println("response header", response.Header)
	})
}
