package integration_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/aaguero96/technical_challenge_q2bank/integration_test/utils"
)

func TestEndpoints(t *testing.T) {
	// Create http client
	client := &http.Client{}

	// EndpointAssert struct
	endpointAssert := utils.EndpointAssert{
		BaseUrl: "http://0.0.0.0:3000",
		T:       t,
		Client:  client,
	}

	// test scenarios
	tests := []struct {
		method   string
		endpoint string
	}{
		{method: "POST", endpoint: "/v1/login"},
		{method: "GET", endpoint: "/v1/register-types"},
		{method: "GET", endpoint: "/v1/transactions"},
		{method: "GET", endpoint: "/v1/transactions/:id"},
		{method: "POST", endpoint: "/v1/transactions/"},
		{method: "DELETE", endpoint: "/v1/transactions/:id"},
		{method: "GET", endpoint: "/v1/user-types"},
		{method: "GET", endpoint: "/v1/user-types/:id"},
		{method: "POST", endpoint: "/v1/users/"},
		{method: "GET", endpoint: "/v1/users"},
		{method: "GET", endpoint: "/v1/users/:id"},
		{method: "PATCH", endpoint: "/v1/wallets/:id"},
		{method: "GET", endpoint: "/v1/wallets"},
		{method: "GET", endpoint: "/v1/wallets/:id"},
	}

	for _, test := range tests {
		title := fmt.Sprintf("Endpoint %v %v exist", test.method, test.endpoint)
		t.Run(title, func(t *testing.T) {
			endpointAssert.Validate(test.method, test.endpoint)
		})
	}
}
