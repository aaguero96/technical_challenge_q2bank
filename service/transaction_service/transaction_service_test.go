package transaction_service

import (
	"testing"

	"github.com/aaguero96/technical_challenge_q2bank/repository/transaction_repository"
	"github.com/aaguero96/technical_challenge_q2bank/repository/user_repository"
	"github.com/aaguero96/technical_challenge_q2bank/repository/user_type_repository"
	"github.com/aaguero96/technical_challenge_q2bank/repository/wallet_repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnitGetAll(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Request / Response
	type Response struct {
		transactions []TransactionResponse
		err          error
	}

	// Mock Repositories
	transactionRepositoryMock := transaction_repository.NewTransactionRepositoryMock()
	userRepositoryMock := user_repository.NewUserRepositoryMock()
	userTypeRepositoryMock := user_type_repository.NewUserTypeRepositoryMock()
	walletRepositoryMock := wallet_repository.NewWalletRepositoryMock()

	// Test scenarios
	tests := []struct {
		title   string
		service TransactionService
		expect  Response
	}{
		{
			title:   "if sucess case, return transactions and nil - OK CASE",
			service: NewTransactionService(transactionRepositoryMock, &userRepositoryMock, &walletRepositoryMock, userTypeRepositoryMock),
			expect: Response{
				transactions: []TransactionResponse{
					{
						PayerID: 1,
						PayeeID: 2,
						Amount:  100,
						Status:  "completed",
					},
				},
				err: nil,
			},
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			result, err := test.service.GetAll()
			if test.expect.err != nil {
				assert.Equal(test.expect.err, err, "expected error %v, instead got %v", test.expect.err, err)
			} else {
				require.Nil(err, "expected error to be nil, instead got %v on %s", err, test.expect.err)
			}
			assert.Equal(test.expect.transactions, result, "expected result %v, instead got %v", test.expect.transactions, result)
		})
	}
}

func TestUnitGetById(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Request / Response
	type Request struct {
		id int
	}
	type Response struct {
		transaction GetByIdResponse
		err         error
	}

	// Mock Repositories
	transactionRepositoryMock := transaction_repository.NewTransactionRepositoryMock()
	userRepositoryMock := user_repository.NewUserRepositoryMock()
	userTypeRepositoryMock := user_type_repository.NewUserTypeRepositoryMock()
	walletRepositoryMock := wallet_repository.NewWalletRepositoryMock()

	// Test scenarios
	tests := []struct {
		title   string
		service TransactionService
		param   Request
		expect  Response
	}{
		{
			title:   "if sucess case, return transaction and nil - OK CASE",
			service: NewTransactionService(transactionRepositoryMock, &userRepositoryMock, &walletRepositoryMock, userTypeRepositoryMock),
			param: Request{
				id: 1,
			},
			expect: Response{
				transaction: GetByIdResponse{
					TransactionID: 1,
					PayerID:       1,
					PayeeID:       2,
					Amount:        100,
					Status:        "completed",
				},
				err: nil,
			},
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			result, err := test.service.GetById(test.param.id)
			if test.expect.err != nil {
				assert.Equal(test.expect.err, err, "expected error %v, instead got %v", test.expect.err, err)
			} else {
				require.Nil(err, "expected error to be nil, instead got %v on %s", err, test.expect.err)
			}
			assert.Equal(test.expect.transaction, result, "expected result %v, instead got %v", test.expect.transaction, result)
		})
	}
}
