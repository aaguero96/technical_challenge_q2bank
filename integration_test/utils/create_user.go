package utils

func CreateUser(name, email, password string, registerNumber, registerTypeID, userTypeID int) string {
	// type Body for these requests
	type Body struct {
		Name           string `json:"name"`
		Email          string `json:"email"`
		Password       string `json:"password"`
		RegisterNumber int    `json:"register_number"`
		RegisterTypeID int    `json:"register_type_id"`
		UserTypeID     int    `json:"user_type_id"`
	}

	// body
	body := Body{
		Name:           name,
		Email:          email,
		Password:       password,
		RegisterNumber: registerNumber,
		RegisterTypeID: registerTypeID,
		UserTypeID:     userTypeID,
	}

	// type reponse
	type responseType struct {
		Token    string `json:"token"`
		ExpingIn string `json:"expiring_in"`
	}

	// response
	_, data, _ := EndpointResponse[responseType](
		EndpointRequest{Body: body}, "POST", "http://0.0.0.0:3000/v1/users?agree_cookie=false",
	)

	return data.Token
}
