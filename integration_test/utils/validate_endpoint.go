package utils

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ValidateEndpoint(t *testing.T, method string, endpoint string) {
	assert := assert.New(t)

	// response
	status, data, _ := EndpointResponse[string](EndpointRequest{Body: nil}, method, endpoint)

	// verify status code
	assert.NotEqual(status, http.StatusNotFound)

	// verify Response
	assert.NotEqual(data, "404 page not found")
}
