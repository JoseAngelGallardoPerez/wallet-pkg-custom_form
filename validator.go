package custom_form

import (
	"reflect"

	"gopkg.in/go-playground/validator.v8"
)

// interface is used to describe validators
type Validator interface {
	// check if validation rules are valid
	Validate(validate *validator.Validate, level *validator.StructLevel, value reflect.Value, fieldName string) *validator.FieldError

	// return all conditions. Validator can be depended from values of another fields
	GetConditions() []*Condition
}
