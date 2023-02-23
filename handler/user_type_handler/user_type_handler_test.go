package user_type_handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aaguero96/technical_challenge_q2bank/repository/user_type_repository"
	"github.com/aaguero96/technical_challenge_q2bank/service/user_type_service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUnitGetAll(t *testing.T) {
	assert := assert.New(t)

	// Intance mock repositoires
	userTypeRepositoryMock := user_type_repository.NewUserTypeRepositoryMock()

	// Instance services
	userTypeService := user_type_service.NewUserTypeService(userTypeRepositoryMock)

	// Request / Response
	type Response struct {
		status int
		body   string
	}

	// Test scenarios
	tests := []struct {
		title   string
		handler UserTypeHandler
		expect  Response
	}{
		{
			title:   "if sucess case, return code status 200 and body response - OK CASE",
			handler: NewUserTypeHandler(userTypeService),

			expect: Response{
				status: http.StatusOK,
				body: func() string {
					jsonBody, _ := json.Marshal([]user_type_service.UserTypeResponse{
						{ID: 1, Type: "common"},
						{ID: 2, Type: "storekeeper"},
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
			router.GET("/user-types", test.handler.GetAll)
			request, _ := http.NewRequest("GET", "/user-types", nil)
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
	userTypeRepositoryMock := user_type_repository.NewUserTypeRepositoryMock()

	// Instance services
	userTypeService := user_type_service.NewUserTypeService(userTypeRepositoryMock)

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
		handler UserTypeHandler
		params  Request
		expect  Response
	}{
		{
			title:   "if sucess case, return code status 200 and body response - OK CASE",
			handler: NewUserTypeHandler(userTypeService),
			params: Request{
				idParam: 1,
			},
			expect: Response{
				status: http.StatusOK,
				body: func() string {
					jsonBody, _ := json.Marshal(user_type_service.GetByIdResponse{UserTypeID: 1, UserType: "common"})
					return string(jsonBody)
				}(),
			},
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			router := gin.Default()
			router.GET("/user-types/:id", test.handler.GetById)
			request, _ := http.NewRequest("GET", fmt.Sprintf("/user-types/%v", test.params.idParam), nil)
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
