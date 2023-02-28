package utils

func LoginUser(email, password string) string {
	// type Body for these requests
	type Body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// body
	body := Body{
		Email:    email,
		Password: password,
	}

	// type reponse
	type responseType struct {
		Token    string `json:"token"`
		ExpingIn string `json:"expiring_in"`
	}

	// response
	_, data, _ := EndpointResponse[responseType](
		EndpointRequest{Body: body}, "POST", BASE_URL+"/v1/login?agree_cookie=true",
	)

	return data.Token
}
