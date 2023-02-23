package wallet_service

import (
	"testing"

	"github.com/aaguero96/technical_challenge_q2bank/repository/wallet_repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnitGetAll(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Request / Response
	type Response struct {
		wallets []WalletResponse
		err     error
	}

	// Mock Repositories
	walletRepositoryMock := wallet_repository.NewWalletRepositoryMock()

	// Test scenarios
	tests := []struct {
		title   string
		service WalletService
		expect  Response
	}{
		{
			title:   "if sucess case, return wallets and nil - OK CASE",
			service: NewWalletService(&walletRepositoryMock),
			expect: Response{
				wallets: []WalletResponse{
					{ID: 1, Amount: 1000},
					{ID: 2, Amount: 2000},
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
			assert.Equal(test.expect.wallets, result, "expected result %v, instead got %v", test.expect.wallets, result)
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
		wallet GetByIdResponse
		err    error
	}

	// Mock Repositories
	walletRepositoryMock := wallet_repository.NewWalletRepositoryMock()

	// Test scenarios
	tests := []struct {
		title   string
		service WalletService
		param   Request
		expect  Response
	}{
		{
			title:   "if sucess case, return wallet and nil - OK CASE",
			service: NewWalletService(&walletRepositoryMock),
			param: Request{
				id: 1,
			},
			expect: Response{
				wallet: GetByIdResponse{WalletID: 1, Amount: 1000},
				err:    nil,
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
			assert.Equal(test.expect.wallet, result, "expected result %v, instead got %v", test.expect.wallet, result)
		})
	}
}

func TestUnitAddAmount(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Request / Response
	type Request struct {
		walletID       int
		increaseAmount float64
	}
	type Response struct {
		wallet AddAmountResponse
		err    error
	}

	// Mock Repositories
	walletRepositoryMock := wallet_repository.NewWalletRepositoryMock()

	// Test scenarios
	tests := []struct {
		title   string
		service WalletService
		param   Request
		expect  Response
	}{
		{
			title:   "if sucess case, return wallet and nil - OK CASE",
			service: NewWalletService(&walletRepositoryMock),
			param: Request{
				walletID:       1,
				increaseAmount: 234,
			},
			expect: Response{
				wallet: AddAmountResponse{WalletID: 1, Amount: 1234},
				err:    nil,
			},
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			result, err := test.service.AddAmount(test.param.walletID, test.param.increaseAmount)
			if test.expect.err != nil {
				assert.Equal(test.expect.err, err, "expected error %v, instead got %v", test.expect.err, err)
			} else {
				require.Nil(err, "expected error to be nil, instead got %v on %s", err, test.expect.err)
			}
			assert.Equal(test.expect.wallet, result, "expected result %v, instead got %v", test.expect.wallet, result)
		})
	}
}
