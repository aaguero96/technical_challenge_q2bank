package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnitValidateRegisterNumber(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	// Request / Response
	type Request struct {
		registerNumber int64
		registerType   string
	}
	type Response struct {
		err error
	}

	// Test scenarios
	tests := []struct {
		title  string
		params Request
		expect Response
	}{
		{
			title: "if passed CPF and number has 11 digits return nil - OK CASE",
			params: Request{
				registerNumber: 12345678900,
				registerType:   "CPF",
			},
			expect: Response{
				err: nil,
			},
		},
		{
			title: "if passed CPF and number has not 11 digits return error - NOK CASE",
			params: Request{
				registerNumber: 1234567890,
				registerType:   "CPF",
			},
			expect: Response{
				err: errors.New("register number is invalid, that was considering the register type passed"),
			},
		},
		{
			title: "if passed CNPJ and number has 14 digits return nil - OK CASE",
			params: Request{
				registerNumber: 12345678900001,
				registerType:   "CNPJ",
			},
			expect: Response{
				err: nil,
			},
		},
		{
			title: "if passed CNPJ and number has not 14 digits return error - NOK CASE",
			params: Request{
				registerNumber: 1234567890000,
				registerType:   "CNPJ",
			},
			expect: Response{
				err: errors.New("register number is invalid, that was considering the register type passed"),
			},
		},
		{
			title: "if passed incorrect register_type return error - NOK CASE",
			params: Request{
				registerNumber: 1234567890000,
				registerType:   "PHONE",
			},
			expect: Response{
				err: errors.New("register_type is incorrect"),
			},
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			err := ValidateRegisterNumber(test.params.registerNumber, test.params.registerType)
			if test.expect.err != nil {
				assert.Equal(test.expect.err, err, "expected error %v, instead got %v", test.expect.err, err)
			} else {
				require.Nil(err, "expected error to be nil, instead got %v on %s", err, test.expect.err)
			}
		})
	}
}
