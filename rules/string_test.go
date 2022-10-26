package rules

import (
	"testing"

	"github.com/jdtron/validate"
	"github.com/stretchr/testify/assert"
)

func TestStringEmpty(t *testing.T) {
	tt := []struct {
		name     string
		value    string
		expectOK bool
	}{
		{name: "empty", value: "", expectOK: true},
		{name: "not empty", value: "not empty", expectOK: false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ok := validate.New().
				Rules(validate.ValidatorMap{
					tc.name: validate.ForField(Str(tc.value).Empty()),
				}).
				Run().
				OK()

			assert.Equal(t, tc.expectOK, ok)
		})
	}
}

func TestStringNotEmpty(t *testing.T) {
	tt := []struct {
		name     string
		value    string
		expectOK bool
	}{
		{name: "empty", value: "", expectOK: false},
		{name: "valid", value: "foobarbaz", expectOK: true},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ok := validate.New().
				Rules(validate.ValidatorMap{
					"foo": validate.ForField(Str(tc.value).NotEmpty()),
				}).
				Run().
				OK()

			assert.Equal(t, tc.expectOK, ok)
		})
	}
}

func TestStringMinLen(t *testing.T) {
	tt := []struct {
		name     string
		min      int
		value    string
		expectOK bool
	}{
		{name: "too short", min: 10, value: "a", expectOK: false},
		{name: "valid", min: 5, value: "aaaaaa", expectOK: true},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ok := validate.New().
				Rules(validate.ValidatorMap{
					"foo": validate.ForField(Str(tc.value).MinLen(tc.min)),
				}).
				Run().
				OK()

			assert.Equal(t, tc.expectOK, ok)
		})
	}
}

func TestStringMaxLen(t *testing.T) {
	tt := []struct {
		name     string
		max      int
		value    string
		expectOK bool
	}{
		{name: "valid", max: 50, value: "foobar", expectOK: true},
		{name: "too long", max: 5, value: "aaaaaaaa", expectOK: false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ok := validate.New().
				Rules(validate.ValidatorMap{
					"foo": validate.ForField(Str(tc.value).MaxLen(tc.max)),
				}).
				Run().
				OK()

			assert.Equal(t, tc.expectOK, ok)
		})
	}
}
