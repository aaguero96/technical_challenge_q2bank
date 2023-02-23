package validator

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type validatorExternalAPI struct{}

func NewValidatorExternalAPI() validatorExternalAPI {
	return validatorExternalAPI{}
}

func (vea validatorExternalAPI) Validation() (bool, error) {
	response, err := http.Get("https://run.mocky.io/v3/d02168c6-d88d-4ff2-aac6-9e9eb3425e31")
	if err != nil {
		return false, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err
	}
	type JsonResponse struct {
		Authorization bool `json:"authorization"`
	}
	var data JsonResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return false, err
	}
	return data.Authorization, nil
}
