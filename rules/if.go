package rules

import "github.com/jdtron/validate"

func If[T any](condition bool, factory func() validate.FieldValidator[T]) validate.FieldValidator[T] {
	if condition {
		return factory()
	}

	return nil
}
