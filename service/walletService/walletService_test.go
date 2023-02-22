package walletService

import (
	"testing"

	"github.com/aaguero96/technical_challenge_q2bank/repository/walletRepository"
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
	walletRepositoryMock := walletRepository.NewWalletRepositoryMock()

	// Test scenarios
	tests := []struct {
		title   string
		service WalletService
		expect  Response
	}{
		{
			title:   "if sucess case, return wallets and nil - OK CASE",
			service: NewWalletService(walletRepositoryMock),
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
