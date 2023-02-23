package register_type_service

import (
	"testing"

	"github.com/aaguero96/technical_challenge_q2bank/repository/register_type_repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnitGetAll(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Request / Response
	type Response struct {
		registerTypes []RegisterTypeResponse
		err           error
	}

	// Mock Repositories
	registerTypeRepositoryMock := register_type_repository.NewRegisterTypeRepositoryMock()

	// Test scenarios
	tests := []struct {
		title   string
		service RegisterTypeService
		expect  Response
	}{
		{
			title:   "if sucess case, return registerTypes and nil - OK CASE",
			service: NewRegisterTypeService(registerTypeRepositoryMock),
			expect: Response{
				registerTypes: []RegisterTypeResponse{
					{ID: 1, Type: "CPF"},
					{ID: 2, Type: "CNPJ"},
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
			assert.Equal(test.expect.registerTypes, result, "expected result %v, instead got %v", test.expect.registerTypes, result)
		})
	}
}
