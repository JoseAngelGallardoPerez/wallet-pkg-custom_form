package custom_form

import (
	"reflect"
	"strconv"

	"github.com/Confialink/wallet-pkg-types"
	"gopkg.in/go-playground/validator.v8"
)

const ValidatorMinName = "min"

type ValidatorMin struct {
	Min        int            `json:"-"`
	Conditions []*Condition   `json:"conditions"`
	Options    types.DataJSON `json:"options"`
	Name       string         `json:"name"`
}

// create new Min validator
func NewValidatorMin(min int, conditions []*Condition) Validator {
	options := make(types.DataJSON)
	options.Set("value", min)
	return &ValidatorMin{Min: min, Conditions: conditions, Options: options, Name: ValidatorMinName}
}

// check if rule is valid
func (v *ValidatorMin) Validate(validate *validator.Validate, level *validator.StructLevel, value reflect.Value, fieldName string) *validator.FieldError {
	valueAsString := strconv.Itoa(v.Min)
	res := validator.HasMinOf(
		validate,
		level.TopStruct,
		level.CurrentStruct,
		value,
		value.Type(),
		value.Kind(),
		valueAsString,
	)

	if !res {
		return makeFieldError(fieldName, value, v.Name, valueAsString)
	}

	return nil
}

func (v *ValidatorMin) GetConditions() []*Condition {
	return v.Conditions
}
