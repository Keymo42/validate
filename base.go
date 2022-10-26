package validate

// ValidatorBase is used to construct a new FieldValidator.
// It contains basic attribues for field value and rules and provides helper methods.
//
// Satisfies validate.FieldValidator[T]
type ValidatorBase[T any] struct {
	Value T
	rules Rules
}

// AddRule appends a rule to the rule stack
func (v *ValidatorBase[T]) AddRule(r Rule) {
	v.rules = append(v.rules, r)
}

// Rules returns all configured validation rules
func (v *ValidatorBase[T]) Rules() Rules {
	return v.rules
}
