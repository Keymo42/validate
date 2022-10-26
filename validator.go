package validate

import "fmt"

// ValidationError represents an error that happened during validation.
//
// Msg is used for exceptional invalid values, like len(str) should be > 42, but is 12.
// Err is used for actual errors that happen during validation, that could not be related to invalid values.
type ValidationError struct {
	Msg string
	Err error
}

// String displays the validation error as a string
func (v *ValidationError) String() string {
	if v.Err != nil {
		return fmt.Sprintf("Validation error: %s", v.Err.Error())
	}

	return fmt.Sprintf("Validation failed: %s", v.Err.Error())
}

// Rule contains actual functionality to validate a value
type Rule func() *ValidationError

// Rules is a collection of Rules
type Rules []Rule

// ValidatorMap maps field names to validation rules
type ValidatorMap map[string]FieldValidator[any]

// ErrorMap maps validation erros to fieldnames
type ErrorMap map[string][]ValidationError

// Add appends a validation error to the field's error stack
func (e *ErrorMap) Add(fieldName string, err ValidationError) {
	(*e)[fieldName] = append((*e)[fieldName], err)
}

// RuleProvider contains methods for providing validation rules
type RuleProvider interface {
	Rules() Rules
}

// FieldValidator provides validation rules
type FieldValidator[T any] interface {
	Rules() Rules
}

// ValidatorOptions contains fields to configure a validator
type ValidatorOptions struct {
	Bail bool
}

// Validator is the central struct used to validate anything.
// It holds the validation map, error stacks and configuration used for validation.
type Validator struct {
	validatorMap ValidatorMap
	errs         ErrorMap
	opts         ValidatorOptions
}

// New create a new Validator
func New() *Validator {
	return &Validator{
		validatorMap: make(ValidatorMap),
		errs:         make(ErrorMap),
	}
}

// Options add validation options
func (v *Validator) Options(opts ValidatorOptions) *Validator {
	v.opts = opts
	return v
}

// Rules set validation rules
func (v *Validator) Rules(rules ValidatorMap) *Validator {
	for fieldName, fieldValidator := range rules {
		v.validatorMap[fieldName] = fieldValidator
	}

	return v
}

// Run validation rules provided with v.Rules under respect of the set validation options
func (v *Validator) Run() *Validator {
	for field, validator := range v.validatorMap {
		for _, rule := range validator.Rules() {
			if err := rule(); err != nil {
				v.errs.Add(field, *err)

				if v.opts.Bail {
					break
				}
			}
		}
	}

	return v
}

// Errs returns validation errors, if any
func (v *Validator) Errs() ErrorMap {
	return v.errs
}

// OK returns true if no valiation errors occured
func (v *Validator) OK() bool {
	return len(v.errs) == 0
}
