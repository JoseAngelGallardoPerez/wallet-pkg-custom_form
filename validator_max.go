package custom_form

import (
	"reflect"
	"strconv"

	"github.com/Confialink/wallet-pkg-types"
	"gopkg.in/go-playground/validator.v8"
)

const ValidatorMaxName = "max"

type ValidatorMax struct {
	Max        int            `json:"-"`
	Conditions []*Condition   `json:"conditions"`
	Options    types.DataJSON `json:"options"`
	Name       string         `json:"name"`
}

// create new Max validator
func NewValidatorMax(max int, conditions []*Condition) Validator {
	options := make(types.DataJSON)
	options.Set("value", max)
	return &ValidatorMax{Max: max, Conditions: conditions, Options: options, Name: ValidatorMaxName}
}

// check if rule is valid
func (v *ValidatorMax) Validate(validate *validator.Validate, structLevel *validator.StructLevel, value reflect.Value, fieldName string) *validator.FieldError {
	valueAsString := strconv.Itoa(v.Max)
	res := validator.HasMaxOf(
		validate,
		structLevel.TopStruct,
		structLevel.CurrentStruct,
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

func (v *ValidatorMax) GetConditions() []*Condition {
	return v.Conditions
}
