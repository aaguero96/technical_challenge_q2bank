package wallet_handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aaguero96/technical_challenge_q2bank/repository/wallet_repository"
	"github.com/aaguero96/technical_challenge_q2bank/service/wallet_service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUnitGetAll(t *testing.T) {
	assert := assert.New(t)

	// Intance mock repositoires
	walletRepositoryMock := wallet_repository.NewWalletRepositoryMock()

	// Instance services
	walletService := wallet_service.NewWalletService(&walletRepositoryMock)

	// Request / Response
	type Response struct {
		status int
		body   string
	}

	// Test scenarios
	tests := []struct {
		title   string
		handler WalletHandler
		expect  Response
	}{
		{
			title:   "if sucess case, return code status 200 and body response - OK CASE",
			handler: NewWalletHandler(walletService),

			expect: Response{
				status: http.StatusOK,
				body: func() string {
					jsonBody, _ := json.Marshal([]wallet_service.WalletResponse{
						{ID: 1, Amount: 1000},
						{ID: 2, Amount: 2000},
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
			router.GET("/wallets", test.handler.GetAll)
			request, _ := http.NewRequest("GET", "/wallets", nil)
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
	walletRepositoryMock := wallet_repository.NewWalletRepositoryMock()

	// Instance services
	walletService := wallet_service.NewWalletService(&walletRepositoryMock)

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
		handler WalletHandler
		params  Request
		expect  Response
	}{
		{
			title:   "if sucess case, return code status 200 and body response - OK CASE",
			handler: NewWalletHandler(walletService),
			params: Request{
				idParam: 1,
			},
			expect: Response{
				status: http.StatusOK,
				body: func() string {
					jsonBody, _ := json.Marshal(wallet_service.GetByIdResponse{WalletID: 1, Amount: 1000})
					return string(jsonBody)
				}(),
			},
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			router := gin.Default()
			router.GET("/wallets/:id", test.handler.GetById)
			request, _ := http.NewRequest("GET", fmt.Sprintf("/wallets/%v", test.params.idParam), nil)
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
