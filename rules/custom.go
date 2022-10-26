package rules

import "github.com/jdtron/validate"

// CustomRule represents a custom validation rule.
// This could either be a predefined function, or a clusore constructed on the go.
type CustomRule[T any] func(value T) *validate.ValidationError

// CustomValidator is a FieldValidator for custom validation rules
type CustomValidator[T any] struct {
	validate.ValidatorBase[T]
}

// Custom creates a new CustomValidator
func Custom[T any](val T, rule CustomRule[T]) *CustomValidator[T] {
	v := CustomValidator[T]{
		validate.ValidatorBase[T]{
			Value: val,
		},
	}

	v.AddRule(func() *validate.ValidationError {
		return rule(v.Value)
	})

	return &v
}
