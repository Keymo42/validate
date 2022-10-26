package rules

import (
	"fmt"

	"github.com/jdtron/validate"
)

// StringValidator provides validation methods for strings
type StringValidator struct {
	validate.ValidatorBase[string]
}

// Str creates a new StringValidator
func Str(value string) *StringValidator {
	return &StringValidator{
		validate.ValidatorBase[string]{
			Value: value,
		},
	}
}

// Empty valdiates that a string must be empty
func (s *StringValidator) Empty() *StringValidator {
	s.AddRule(func() *validate.ValidationError {
		if s.Value != "" {
			return &validate.ValidationError{
				Msg: "must be empty",
			}
		}

		return nil
	})

	return s
}

// NotEmpty validates that a string must be not empty
func (s *StringValidator) NotEmpty() *StringValidator {
	s.AddRule(func() *validate.ValidationError {
		if s.Value == "" {
			return &validate.ValidationError{
				Msg: "must not be empty",
			}
		}

		return nil
	})

	return s
}

// NotEmpty validates that a string must be at least n characters long
func (s *StringValidator) MinLen(min int) *StringValidator {
	s.AddRule(func() *validate.ValidationError {
		if len(s.Value) < min {
			return &validate.ValidationError{
				Msg: fmt.Sprintf("must be at least %d characters long, but is %d", min, len(s.Value)),
			}
		}

		return nil
	})

	return s
}

// MaxLen validates that a string must be at most n characters long
func (s *StringValidator) MaxLen(max int) *StringValidator {
	s.AddRule(func() *validate.ValidationError {
		if len(s.Value) > max {
			return &validate.ValidationError{
				Msg: fmt.Sprintf("must be at most %d characters long, but is %d", max, len(s.Value)),
			}
		}

		return nil
	})

	return s
}
