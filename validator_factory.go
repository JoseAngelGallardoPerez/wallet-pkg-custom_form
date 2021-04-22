package custom_form

import (
	"errors"

	"github.com/Confialink/wallet-pkg-types"
)

// interface to create a new validator
type ValidatorFactoryInterface interface {
	// make a new Validator
	Make(validatorName string, conditions []*Condition, options types.DataJSON) (Validator, error)

	// list of available validators which the factory can create
	AvailableValidators() map[string]bool
}

// validator factory struct
type ValidatorFactory struct {
	factories []ValidatorFactoryInterface
}

func NewValidatorFactory() *ValidatorFactory {
	factory := &ValidatorFactory{}
	factory.Register(&DefaultValidatorFactory{})
	return factory
}

// register custom validator factory
func (v *ValidatorFactory) Register(factory ValidatorFactoryInterface) {
	v.factories = append(v.factories, factory)
}

// make validator. make default validators or from custom registered factory
func (v *ValidatorFactory) Make(validatorName string, conditions []*Condition, options types.DataJSON) (Validator, error) {
	for _, factory := range v.factories {
		if validator, _ := factory.Make(validatorName, conditions, options); validator != nil {
			return validator, nil
		}
	}

	return nil, errors.New("validator not found in factories")
}

// return list of all registered validators
func (v *ValidatorFactory) AvailableValidators() map[string]bool {
	availableValidators := make(map[string]bool)
	for _, factory := range v.factories {
		for key, value := range factory.AvailableValidators() {
			availableValidators[key] = value
		}
	}

	return availableValidators
}

// default validator factory struct
type DefaultValidatorFactory struct {
}

func (v *DefaultValidatorFactory) Make(validatorName string, conditions []*Condition, options types.DataJSON) (Validator, error) {
	switch validatorName {
	case ValidatorMaxName:
		return NewValidatorMax(int(options.GetFloat64("value")), conditions), nil
	case ValidatorMinName:
		return NewValidatorMin(int(options.GetFloat64("value")), conditions), nil
	}

	return nil, errors.New("validator not found in factories")
}

func (v *DefaultValidatorFactory) AvailableValidators() map[string]bool {
	return map[string]bool{
		ValidatorMaxName: true,
		ValidatorMinName: true,
	}
}
