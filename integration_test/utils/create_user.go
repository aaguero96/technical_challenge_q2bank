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
		Name:           "name_0",
		Email:          "name0@test.com",
		Password:       "Def4!t*0",
		RegisterNumber: 12345678900,
		RegisterTypeID: 1,
		UserTypeID:     1,
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
