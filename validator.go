package validate

import "fmt"

// ValidationError represents an error that happened during validation.
//
// Msg is used for exceptional invalid values, like len(str) should be > 42, but is 12.
// Err is used for actual errors that happen during validation, that could not be related to invalid values.
type ValidationError struct {
	Msg  string
	Err  error
	Code int
}

// String displays the validation error as a string
func (v *ValidationError) String() string {
	var str string

	if v.Err != nil {
		str = fmt.Sprintf("Validation error: %s", v.Err.Error())
	} else {
		str = fmt.Sprintf("Validation failed: %s", v.Msg)
	}

	if v.Code != 0 {
		str = fmt.Sprintf("%s; code=%d", str, v.Code)
	}

	return str
}

// applyDefaults appies default values specified in ValidatorOptions
func (v *ValidationError) applyDefaults(opts *ValidatorOptions) {
	if v.Code == 0 && opts.DefaultCode != 0 {
		v.Code = opts.DefaultCode
	}
}

// Rule contains actual functionality to validate a value
type Rule func() *ValidationError

// Rules is a collection of Rules
type Rules []Rule

// ValidatorMap maps field names to validation rules
type ValidatorMap map[string][]FieldValidator[any]

// ForField is a wrapper function that helps you define lists of FieldValidators
func ForField(v ...FieldValidator[any]) []FieldValidator[any] {
	return v
}

// ErrorMap maps validation erros to fieldnames
type ErrorMap map[string][]ValidationError

// First returns the first affected field name and it's corresponding validation error, if any.
func (e ErrorMap) First() (string, ValidationError) {
	unknownErr := ValidationError{
		Msg: "Unknown error",
	}

	for fieldName, errs := range e {
		var err ValidationError
		if len(errs) == 0 {
			err = unknownErr
		} else {
			err = errs[0]
		}

		return fieldName, err
	}

	return "", unknownErr
}

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
	Bail        bool
	DefaultCode int
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
	for field, validators := range v.validatorMap {
		defer v.recoverRulePanic(field)

		for _, validator := range validators {
			for _, rule := range validator.Rules() {
				if err := rule(); err != nil {
					err.applyDefaults(&v.opts)
					v.errs.Add(field, *err)

					if v.opts.Bail {
						break
					}
				}
			}
		}
	}

	return v
}

// With is easy and quick method to add a validators to a field
func (v *Validator) With(fieldName string, rules ...FieldValidator[any]) *Validator {
	v.validatorMap[fieldName] = append(v.validatorMap[fieldName], rules...)

	return v
}

// WithRecover is the same as With, but recovers from errors
func (v *Validator) WithRecover(fieldName string, closure func(validator *Validator)) *Validator {
	defer (func() {
		if r := recover(); r != nil {
			v.errs.Add(fieldName, ValidationError{
				Err: fmt.Errorf("panic while adding rule: %#v", r),
			})
		}
	})()

	closure(v)
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

// recoverRulePanic recovers a panics that happen during validation
func (v *Validator) recoverRulePanic(fieldName string) {
	if r := recover(); r != nil {
		v.errs.Add(fieldName, ValidationError{
			Err: fmt.Errorf("panic while validating: %#v", r),
		})
	}
}
