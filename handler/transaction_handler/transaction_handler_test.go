package transaction_handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aaguero96/technical_challenge_q2bank/repository/transaction_repository"
	"github.com/aaguero96/technical_challenge_q2bank/repository/user_repository"
	"github.com/aaguero96/technical_challenge_q2bank/repository/user_type_repository"
	"github.com/aaguero96/technical_challenge_q2bank/repository/wallet_repository"
	"github.com/aaguero96/technical_challenge_q2bank/service/transaction_service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUnitGetAll(t *testing.T) {
	assert := assert.New(t)

	// Intance mock repositoires
	transactionRepositoryMock := transaction_repository.NewTransactionRepositoryMock()
	userRepositoryMock := user_repository.NewUserRepositoryMock()
	walletRepositoryMock := wallet_repository.NewWalletRepositoryMock()
	user_typeRepositoryMock := user_type_repository.NewUserTypeRepositoryMock()

	// Instance services
	transactionService := transaction_service.NewTransactionService(&transactionRepositoryMock, &userRepositoryMock, &walletRepositoryMock, user_typeRepositoryMock)

	// Request / Response
	type Response struct {
		status int
		body   string
	}

	// Test scenarios
	tests := []struct {
		title   string
		handler TransactionHandler
		expect  Response
	}{
		{
			title:   "if sucess case, return code status 200 and body response - OK CASE",
			handler: NewTransactionHandler(transactionService),

			expect: Response{
				status: http.StatusOK,
				body: func() string {
					jsonBody, _ := json.Marshal([]transaction_service.TransactionResponse{
						{PayerID: 1, PayeeID: 2, Amount: 100, Status: "completed"},
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
			router.GET("/trasactions", test.handler.GetAll)
			request, _ := http.NewRequest("GET", "/trasactions", nil)
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
	transactionRepositoryMock := transaction_repository.NewTransactionRepositoryMock()
	userRepositoryMock := user_repository.NewUserRepositoryMock()
	walletRepositoryMock := wallet_repository.NewWalletRepositoryMock()
	user_typeRepositoryMock := user_type_repository.NewUserTypeRepositoryMock()

	// Instance services
	transactionService := transaction_service.NewTransactionService(&transactionRepositoryMock, &userRepositoryMock, &walletRepositoryMock, user_typeRepositoryMock)

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
		handler TransactionHandler
		params  Request
		expect  Response
	}{
		{
			title:   "if sucess case, return code status 200 and body response - OK CASE",
			handler: NewTransactionHandler(transactionService),
			params: Request{
				idParam: 1,
			},
			expect: Response{
				status: http.StatusOK,
				body: func() string {
					jsonBody, _ := json.Marshal(transaction_service.GetByIdResponse{
						TransactionID: 1,
						PayerID:       1,
						PayeeID:       2,
						Amount:        100,
						Status:        "completed",
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
			router.GET("/trasactions/:id", test.handler.GetById)
			request, _ := http.NewRequest("GET", fmt.Sprintf("/trasactions/%v", test.params.idParam), nil)
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
