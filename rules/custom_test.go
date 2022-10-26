package rules

import (
	"fmt"
	"testing"

	"github.com/jdtron/validate"
	"github.com/stretchr/testify/assert"
)

func validateIntIs42(val int) *validate.ValidationError {
	if val != 42 {
		return &validate.ValidationError{
			Msg: "must equal 42",
		}
	}

	return nil
}

func validateStrLenIsBetween10And20(val string) *validate.ValidationError {
	if len(val) < 10 || len(val) > 20 {
		return &validate.ValidationError{
			Msg: fmt.Sprintf("len must be between 10 and 20 but is %d", len(val)),
		}
	}

	return nil
}

func TestCustomInt(t *testing.T) {
	tt := []struct {
		name     string
		value    int
		expectOK bool
	}{
		{name: "12", value: 12, expectOK: false},
		{name: "42", value: 42, expectOK: true},
		{name: "-99", value: -99, expectOK: false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ok := validate.New().
				Rules(validate.ValidatorMap{
					"foo": Custom(tc.value, validateIntIs42),
				}).
				Run().
				OK()
			assert.Equal(t, tc.expectOK, ok)
		})
	}
}

func TestCustomString(t *testing.T) {
	tt := []struct {
		name     string
		value    string
		expectOK bool
	}{
		{name: "too short", value: "a", expectOK: false},
		{name: "valid", value: "aaaaaaaaaaaaaaa", expectOK: true},
		{name: "too long", value: "aaaaaaaaaaaaaaaaaaaaa", expectOK: false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ok := validate.New().
				Rules(validate.ValidatorMap{
					"foo": Custom(tc.value, validateStrLenIsBetween10And20),
				}).
				Run().
				OK()
			assert.Equal(t, tc.expectOK, ok)
		})
	}
}
