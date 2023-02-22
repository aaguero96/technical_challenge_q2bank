package externalValidatorService

import (
	"testing"

	"github.com/aaguero96/technical_challenge_q2bank/externalAPI/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnitValidator(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Request / Response
	type Response struct {
		authorization bool
		err           error
	}

	// Test scenarios
	tests := []struct {
		title   string
		service ExternalValidatorService
		expect  Response
	}{
		{
			title:   "if sucess case, return authorization and nil - OK CASE",
			service: NewExternalValidatorService(validator.NewValidatorExternalAPIMock(true)),
			expect: Response{
				authorization: true,
				err:           nil,
			},
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			result, err := test.service.Validator()
			if test.expect.err != nil {
				assert.Equal(test.expect.err, err, "expected error %v, instead got %v", test.expect.err, err)
			} else {
				require.Nil(err, "expected error to be nil, instead got %v on %s", err, test.expect.err)
			}
			assert.Equal(test.expect.authorization, result, "expected result %v, instead got %v", test.expect.authorization, result)
		})
	}
}
