package tests

import (
	"errors"
	"testing"

	"github.com/jdtron/validate"
	"github.com/jdtron/validate/rules"
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
		err       validate.ValidationError
	}{
		{name: "with msg", fieldName: "field1", err: validate.ValidationError{Msg: "Something went wrong"}},
		{name: "with err", fieldName: "field2", err: validate.ValidationError{Err: errors.New("Oops!")}},
	}

	for _, tc := range tt {
		s.T().Run(tc.name, func(t *testing.T) {
			em := make(validate.ErrorMap)
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
	opts := validate.ValidatorOptions{
		DefaultCode: 42,
	}

	err := validate.ValidationError{Msg: "Oopsie"}
	err.ApplyDefaults(&opts)
	assert.Equal(v.T(), opts.DefaultCode, err.Code)
}

func (v *ValidationErrorSuite) TestBailOption() {
	t := v.T()

	type TestStruct struct {
		innerStrct *TestStruct
	}

	cases := []struct {
		name      string
		strct     *TestStruct
		expectErr bool
	}{
		{
			name: "validation success",
			strct: &TestStruct{
				innerStrct: &TestStruct{},
			},
			expectErr: false,
		},
		{
			name:      "validation fail",
			strct:     &TestStruct{},
			expectErr: true,
		},
	}

	for _, c := range cases {
		validator := validate.New().
			Rules(validate.ValidatorMap{
				"innerStrct": validate.ForField(
					rules.Any(c.strct.innerStrct).NotNil(),
				),
			})

		ok := validator.Run().OK()
		field, _ := validator.Errs().First()

		if c.expectErr {
			assert.False(t, ok)
			assert.Equal(t, "innerStrct", field)
		} else {
			assert.True(t, ok)
		}
	}
}

func TestErrorMapSuite(t *testing.T) {
	suite.Run(t, new(ErrorMapSuite))
}

func TestValidationErrorSuite(t *testing.T) {
	suite.Run(t, new(ValidationErrorSuite))
}
