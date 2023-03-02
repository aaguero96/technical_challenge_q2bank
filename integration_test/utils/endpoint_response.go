package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type EndpointHeader struct {
	BearerToken string
}

type EndpointRequest struct {
	Body   interface{}
	Header EndpointHeader
}

func EndpointResponse[T any](request EndpointRequest, method string, endpoint string) (int, T, []*http.Cookie) {
	// http client
	client := &http.Client{}

	// body
	body, err := json.Marshal(request.Body)
	if err != nil {
		fmt.Println("error when parsisng json")
		fmt.Println(err)
	}

	// http request
	httpRequest, err := http.NewRequest(method, endpoint, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("error when request is created")
		fmt.Println(err)
	}

	// headers
	httpRequest.Header.Add("Authorization", request.Header.BearerToken)

	// http response
	response, err := client.Do(httpRequest)
	if err != nil {
		fmt.Println("error when response is created")
		fmt.Println(err)
	}
	defer response.Body.Close()

	// read response data
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error to read response")
		fmt.Println(err)
	}
	var data T
	if err = json.Unmarshal([]byte(string(responseData)), &data); err != nil {
		fmt.Println("data dont has correct types")
	}

	return response.StatusCode, data, response.Cookies()
}
