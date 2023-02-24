package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type EndpointAssert struct {
	BaseUrl string
	T       *testing.T
	Client  *http.Client
}

func (ea EndpointAssert) Validate(method string, endpoint string) {
	assert := assert.New(ea.T)
	require := require.New(ea.T)

	// http request
	url := fmt.Sprintf("%v%v", ea.BaseUrl, endpoint)
	request, _ := http.NewRequest(method, url, nil)

	// http response
	response, err := ea.Client.Do(request)
	if err != nil {
		require.Nil(err)
	}
	defer response.Body.Close()

	// verify status code
	assert.NotEqual(response.StatusCode, http.StatusNotFound)

	// verify Response
	responseData, _ := ioutil.ReadAll(response.Body)
	var data string
	json.Unmarshal([]byte(string(responseData)), &data)
	assert.NotEqual(data, "404 page not found")
}
