package userTypeService

import (
	"testing"

	"github.com/aaguero96/technical_challenge_q2bank/repository/userTypeRepository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnitGetAll(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Request / Response
	type Response struct {
		userTypes []UserTypeResponse
		err       error
	}

	// Mock Repositories
	userTypeRepositoryMock := userTypeRepository.NewUserTypeRepositoryMock()

	// Test scenarios
	tests := []struct {
		title   string
		service UserTypeService
		expect  Response
	}{
		{
			title:   "if sucess case, return wallets and nil - OK CASE",
			service: NewUserTypeService(userTypeRepositoryMock),
			expect: Response{
				userTypes: []UserTypeResponse{
					{ID: 1, Type: "common"},
					{ID: 2, Type: "storekeeper"},
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
			assert.Equal(test.expect.userTypes, result, "expected result %v, instead got %v", test.expect.userTypes, result)
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
		userType GetByIdResponse
		err      error
	}

	// Mock Repositories
	userTypeRepositoryMock := userTypeRepository.NewUserTypeRepositoryMock()

	// Test scenarios
	tests := []struct {
		title   string
		service userTypeService
		param   Request
		expect  Response
	}{
		{
			title:   "if sucess case, return wallet and nil - OK CASE",
			service: NewUserTypeService(userTypeRepositoryMock),
			param: Request{
				id: 1,
			},
			expect: Response{
				userType: GetByIdResponse{UserTypeID: 1, UserType: "common"},
				err:      nil,
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
			assert.Equal(test.expect.userType, result, "expected result %v, instead got %v", test.expect.userType, result)
		})
	}
}
