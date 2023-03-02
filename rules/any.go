package rules

import (
	"reflect"

	"github.com/jdtron/validate"
)

type AnyValidator struct {
	validate.ValidatorBase[any]
}

// Any creates a new AnyValidator
func Any[T any](value T) *AnyValidator {
	return &AnyValidator{
		validate.ValidatorBase[any]{
			Value: value,
		},
	}
}

// Nil validates that a value must be nil
func (a *AnyValidator) Nil() *AnyValidator {
	a.AddRule(func() *validate.ValidationError {
		if a.Value != nil && !reflect.ValueOf(a.Value).IsNil() {
			return &validate.ValidationError{
				Msg: "must be nil",
			}
		}

		return nil
	})

	return a
}

// NotNil validates that a value must be not nil
func (a *AnyValidator) NotNil() *AnyValidator {
	a.AddRule(func() *validate.ValidationError {
		if a.Value == nil || reflect.ValueOf(a.Value).IsNil() {
			return &validate.ValidationError{
				Msg: "must be not nil",
			}
		}

		return nil
	})

	return a
}
