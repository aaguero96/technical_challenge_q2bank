package user_handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aaguero96/technical_challenge_q2bank/repository/user_repository"
	"github.com/aaguero96/technical_challenge_q2bank/service/user_service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUnitGetAll(t *testing.T) {
	assert := assert.New(t)

	// Intance mock repositoires
	userRepositoryMock := user_repository.NewUserRepositoryMock()

	// Instance services
	userService := user_service.NewUserService(&userRepositoryMock)

	// Request / Response
	type Response struct {
		status int
		body   string
	}

	// Test scenarios
	tests := []struct {
		title   string
		handler UserHandler
		expect  Response
	}{
		{
			title:   "if sucess case, return code status 200 and body response - OK CASE",
			handler: NewUserHandler(userService),

			expect: Response{
				status: http.StatusOK,
				body: func() string {
					jsonBody, _ := json.Marshal([]user_service.UserResponse{
						{Name: "name_1", Email: "email1@test.com", WalletID: 1},
						{Name: "name_2", Email: "email2@test.com", WalletID: 2},
					})
					return string(jsonBody)
				}(),
			},
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			router := gin.Default()
			router.GET("/users", test.handler.GetAll)
			request, _ := http.NewRequest("GET", "/users", nil)
			writter := httptest.NewRecorder()
			router.ServeHTTP(writter, request)

			// Verify status code
			assert.Equal(test.expect.status, writter.Code)

			// Verify response
			responseData, _ := ioutil.ReadAll(writter.Body)
			assert.Equal(test.expect.body, string(responseData), "expected result %v, instead got %v", test.expect.body, responseData)
		})
	}
}

func TestUnitGetById(t *testing.T) {
	assert := assert.New(t)

	// Intance mock repositoires
	userRepositoryMock := user_repository.NewUserRepositoryMock()

	// Instance services
	userService := user_service.NewUserService(&userRepositoryMock)

	// Request / Response
	type Request struct {
		idParam int
	}
	type Response struct {
		status int
		body   string
	}

	// Test scenarios
	tests := []struct {
		title   string
		handler UserHandler
		params  Request
		expect  Response
	}{
		{
			title:   "if sucess case, return code status 200 and body response - OK CASE",
			handler: NewUserHandler(userService),
			params: Request{
				idParam: 1,
			},
			expect: Response{
				status: http.StatusOK,
				body: func() string {
					jsonBody, _ := json.Marshal(user_service.GetByIdResponse{
						Name:           "name_1",
						RegisterNumber: 12345678900,
						RegisterTypeID: 1,
						Email:          "email1@test.com",
						WalletID:       1,
						UserTypeID:     1,
					})
					return string(jsonBody)
				}(),
			},
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			router := gin.Default()
			router.GET("/users/:id", test.handler.GetById)
			request, _ := http.NewRequest("GET", fmt.Sprintf("/users/%v", test.params.idParam), nil)
			writter := httptest.NewRecorder()
			router.ServeHTTP(writter, request)

			// Verify status code
			assert.Equal(test.expect.status, writter.Code)

			// Verify response
			responseData, _ := ioutil.ReadAll(writter.Body)
			assert.Equal(test.expect.body, string(responseData), "expected result %v, instead got %v", test.expect.body, responseData)
		})
	}
}
