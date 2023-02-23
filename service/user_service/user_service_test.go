package user_service

import (
	"testing"

	"github.com/aaguero96/technical_challenge_q2bank/repository/user_repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnitGetAll(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Request / Response
	type Response struct {
		users []UserResponse
		err   error
	}

	// Mock Repositories
	userRepositoryMock := user_repository.NewUserRepositoryMock()

	// Test scenarios
	tests := []struct {
		title   string
		service UserService
		expect  Response
	}{
		{
			title:   "if sucess case, return users and nil - OK CASE",
			service: NewUserService(&userRepositoryMock),
			expect: Response{
				users: []UserResponse{
					{
						Name:     "name_1",
						Email:    "email1@test.com",
						WalletID: 1,
					},
					{
						Name:     "name_2",
						Email:    "email2@test.com",
						WalletID: 2,
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
			assert.Equal(test.expect.users, result, "expected result %v, instead got %v", test.expect.users, result)
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
		user GetByIdResponse
		err  error
	}

	// Mock Repositories
	userRepositoryMock := user_repository.NewUserRepositoryMock()

	// Test scenarios
	tests := []struct {
		title   string
		service userService
		param   Request
		expect  Response
	}{
		{
			title:   "if sucess case, return user and nil - OK CASE",
			service: NewUserService(&userRepositoryMock),
			param: Request{
				id: 1,
			},
			expect: Response{
				user: GetByIdResponse{
					Name:           "name_1",
					RegisterNumber: 12345678900,
					RegisterTypeID: 1,
					Email:          "email1@test.com",
					WalletID:       1,
					UserTypeID:     1,
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
			assert.Equal(test.expect.user, result, "expected result %v, instead got %v", test.expect.user, result)
		})
	}
}

func TestUnitCreateUser(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Request / Response
	type Request struct {
		name           string
		email          string
		password       string
		registerNumber int64
		registerTypeID int
		userTypeID     int
	}
	type Response struct {
		user CreateUserResponse
		err  error
	}

	// Mock Repositories
	userRepositoryMock := user_repository.NewUserRepositoryMock()

	// Test scenarios
	tests := []struct {
		title   string
		service userService
		param   Request
		expect  Response
	}{
		{
			title:   "if sucess case, return user and nil - OK CASE",
			service: NewUserService(&userRepositoryMock),
			param: Request{
				name:           "name_3",
				email:          "email3@teste.com",
				password:       "Def4!t*3",
				registerNumber: 12345678900,
				registerTypeID: 1,
				userTypeID:     1,
			},
			expect: Response{
				user: CreateUserResponse{
					Token:      "",
					ExpiringIn: "30 minutes",
				},
				err: nil,
			},
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			result, err := test.service.CreateUser(test.param.name, test.param.email, test.param.password, test.param.registerNumber, test.param.registerTypeID, test.param.userTypeID)
			if test.expect.err != nil {
				assert.Equal(test.expect.err, err, "expected error %v, instead got %v", test.expect.err, err)
			} else {
				require.Nil(err, "expected error to be nil, instead got %v on %s", err, test.expect.err)
			}
			assert.NotEqual("", result.Token, "token not be empty")
			test.expect.user.Token = result.Token
			assert.Equal(test.expect.user, result, "expected result %v, instead got %v", test.expect.user, result)
		})
	}
}

func TestUnitLoginUser(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Request / Response
	type Request struct {
		email    string
		password string
	}
	type Response struct {
		user LoginUserResponse
		err  error
	}

	// Mock Repositories
	userRepositoryMock := user_repository.NewUserRepositoryMock()

	// Test scenarios
	tests := []struct {
		title   string
		service userService
		param   Request
		expect  Response
	}{
		{
			title:   "if sucess case, return user and nil - OK CASE",
			service: NewUserService(&userRepositoryMock),
			param: Request{
				email:    "email1@teste.com",
				password: "Def4!t*1",
			},
			expect: Response{
				user: LoginUserResponse{
					Token:      "",
					ExpiringIn: "30 minutes",
				},
				err: nil,
			},
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			result, err := test.service.LoginUser(test.param.email, test.param.password)
			if test.expect.err != nil {
				assert.Equal(test.expect.err, err, "expected error %v, instead got %v", test.expect.err, err)
			} else {
				require.Nil(err, "expected error to be nil, instead got %v on %s", err, test.expect.err)
			}
			assert.NotEqual("", result.Token, "token not be empty")
			test.expect.user.Token = result.Token
			assert.Equal(test.expect.user, result, "expected result %v, instead got %v", test.expect.user, result)
		})
	}
}
