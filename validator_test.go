package validate

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ErrorMapSuite struct {
	suite.Suite
}

func (s *ErrorMapSuite) TestAdd() {
	tt := []struct {
		name      string
		fieldName string
		err       ValidationError
	}{
		{name: "with msg", fieldName: "field1", err: ValidationError{Msg: "Something went wrong"}},
		{name: "with err", fieldName: "field2", err: ValidationError{Err: errors.New("Oops!")}},
	}

	for _, tc := range tt {
		s.T().Run(tc.name, func(t *testing.T) {
			em := make(ErrorMap)
			em.Add(tc.fieldName, tc.err)
			assert.Len(s.T(), em, 1)

			fieldName, err := em.First()
			assert.Equal(s.T(), tc.fieldName, fieldName)
			assert.Equal(s.T(), tc.err, err)
		})
	}
}

type ValidationErrorSuite struct {
	suite.Suite
}

func (v *ValidationErrorSuite) TestDefaults() {
	opts := ValidatorOptions{
		DefaultCode: 42,
	}

	err := ValidationError{Msg: "Oopsie"}
	err.applyDefaults(&opts)
	assert.Equal(v.T(), opts.DefaultCode, err.Code)
}

func TestErrorMapSuite(t *testing.T) {
	suite.Run(t, new(ErrorMapSuite))
}

func TestValidationErrorSuite(t *testing.T) {
	suite.Run(t, new(ValidationErrorSuite))
}
