package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Request struct {
	Body interface{}
}

func Response[T any](request Request, method string, endpoint string) (int, T) {
	// http client
	client := &http.Client{}

	// body
	body, err := json.Marshal(request.Body)
	if err != nil {
		fmt.Println("error when parsisng json")
	}

	// http request
	httpRequest, err := http.NewRequest(method, endpoint, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("error when request is created")
	}

	// http response
	response, err := client.Do(httpRequest)
	if err != nil {
		fmt.Println("error when response is created")
	}
	defer response.Body.Close()

	// read response data
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error to read response")
	}
	var data T
	if err = json.Unmarshal([]byte(string(responseData)), &data); err != nil {
		fmt.Println("data dont has correct types")
	}

	return response.StatusCode, data
}
