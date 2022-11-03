package rules

import (
	"errors"
	"testing"

	"github.com/jdtron/validate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AnyTestSuite struct {
	suite.Suite
}

func (a *AnyTestSuite) TestNil() {
	t := a.T()

	tt := []struct {
		name     string
		value    any
		expectOK bool
	}{
		{name: "nil", value: nil, expectOK: true},
		{name: "not-nil", value: "foobar", expectOK: false},
		{name: "err", value: errors.New("an error"), expectOK: false},
	}

	for _, tc := range tt {
		a.Run(tc.name, func() {
			ok := validate.New().
				Rules(validate.ValidatorMap{
					tc.name: validate.ForField(Any(tc.value).Nil()),
				}).
				Run().
				OK()
			assert.Equal(t, tc.expectOK, ok)
		})
	}
}

func TestAnyTestSuite(t *testing.T) {
	suite.Run(t, new(AnyTestSuite))
}
