package integration_test

import (
	"fmt"
	"testing"

	"github.com/aaguero96/technical_challenge_q2bank/integration_test/utils"
)

func TestEndpoints(t *testing.T) {
	// BaseURL
	baseURL := "http://0.0.0.0:3000"

	// test scenarios
	tests := []struct {
		method   string
		endpoint string
	}{
		{method: "POST", endpoint: baseURL + "/v1/login"},
		{method: "GET", endpoint: baseURL + "/v1/register-types"},
		{method: "GET", endpoint: baseURL + "/v1/transactions"},
		{method: "GET", endpoint: baseURL + "/v1/transactions/:id"},
		{method: "POST", endpoint: baseURL + "/v1/transactions/"},
		{method: "DELETE", endpoint: baseURL + "/v1/transactions/:id"},
		{method: "GET", endpoint: baseURL + "/v1/user-types"},
		{method: "GET", endpoint: baseURL + "/v1/user-types/:id"},
		{method: "POST", endpoint: baseURL + "/v1/users/"},
		{method: "GET", endpoint: baseURL + "/v1/users"},
		{method: "GET", endpoint: baseURL + "/v1/users/:id"},
		{method: "PATCH", endpoint: baseURL + "/v1/wallets/:id"},
		{method: "GET", endpoint: baseURL + "/v1/wallets"},
		{method: "GET", endpoint: baseURL + "/v1/wallets/:id"},
	}

	for _, test := range tests {
		title := fmt.Sprintf("Endpoint %v %v exist", test.method, test.endpoint)
		t.Run(title, func(t *testing.T) {
			utils.ValidateEndpoint(t, test.method, test.endpoint)
		})
	}
}
