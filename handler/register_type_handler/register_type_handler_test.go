package register_type_handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aaguero96/technical_challenge_q2bank/repository/register_type_repository"
	"github.com/aaguero96/technical_challenge_q2bank/service/register_type_service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUnitGetAll(t *testing.T) {
	assert := assert.New(t)

	// Intance mock repositoires
	registerTypeRepositoryMock := register_type_repository.NewRegisterTypeRepositoryMock()

	// Instance services
	registerTypeService := register_type_service.NewRegisterTypeService(registerTypeRepositoryMock)

	// Request / Response
	type Response struct {
		status int
		body   string
	}

	// Test scenarios
	tests := []struct {
		title   string
		handler RegisterTypeHandler
		expect  Response
	}{
		{
			title:   "if sucess case, return code status 200 and body response - OK CASE",
			handler: NewRegisterTypeHandler(registerTypeService),

			expect: Response{
				status: http.StatusOK,
				body: func() string {
					jsonBody, _ := json.Marshal([]register_type_service.RegisterTypeResponse{
						{ID: 1, Type: "CPF"},
						{ID: 2, Type: "CNPJ"},
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
			router.GET("/register-types", test.handler.GetAll)
			request, _ := http.NewRequest("GET", "/register-types", nil)
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
