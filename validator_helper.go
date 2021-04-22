package custom_form

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	validatorPkg "gopkg.in/go-playground/validator.v8"
)

type Helper struct {
}

// make validation callback for custom validation.
// we build validation rules by validators in form fields
func (s *Helper) MakeValidationCallback() validatorPkg.StructLevelFunc {
	return func(v *validatorPkg.Validate, sl *validatorPkg.StructLevel) {
		formContainer, ok := sl.CurrentStruct.Interface().(FormContainer)
		if !ok {
			log.Printf("the struct %s does not impement FormContainer interface", sl.CurrentStruct.Type().String())
			return
		}

		for _, field := range formContainer.GetCustomForm().Fields {
			fieldValue := reflect.Indirect(sl.CurrentStruct).FieldByName(strings.Title(field.Name))
			for _, val := range field.Validators {
				if s.isValidConditions(val.GetConditions(), v, sl) {
					if err := val.Validate(v, sl, fieldValue, field.Name); err != nil {
						validationErrors := validatorPkg.ValidationErrors{}
						relativeKey := "form" // todo: change it if need. now it is random
						validationErrors[relativeKey] = err
						sl.ReportValidationErrors(field.Name, validationErrors)
					}
				}
			}
		}
	}
}

// Validator can be depended from conditions. if conditions are valid we can call main validator
func (s *Helper) isValidConditions(conditions []*Condition, validate *validatorPkg.Validate, level *validatorPkg.StructLevel) bool {
	if len(conditions) == 0 {
		return true
	}

	for _, dep := range conditions {
		fieldValue := reflect.Indirect(level.CurrentStruct).FieldByName(strings.Title(dep.FieldName))
		var tagValue string

		// todo: implement other conditions
		if dep.Type == ConditionTypeIn {
			for i, value := range dep.Values {
				if i != 0 {
					tagValue += "|"
				}
				tagValue += "eq=" + fmt.Sprintf("%v", value)
			}
		}

		if len(tagValue) > 0 {
			err := validate.Field(fieldValue.Interface(), tagValue)
			if err != nil {
				return false
			}
		}
	}

	return true
}

// make validation error for field
func makeFieldError(fieldName string, fieldValue reflect.Value, tag string, param string) *validatorPkg.FieldError {
	return &validatorPkg.FieldError{
		FieldNamespace: fieldName,
		NameNamespace:  fieldName,
		Name:           fieldName,
		Field:          fieldName,
		Tag:            tag,
		ActualTag:      tag,
		Param:          param,
		Value:          fieldValue.Interface(),
		Type:           fieldValue.Type(),
	}
}
