## Package provides simple way to validate structures 

### Description  
You can create a form with fields an its validators and validate a struct. It can be useful if the struct has dynamic validation rules

### How to use it?  
```go
package x

import (
	"github.com/Confialink/wallet-pkg-custom_form"
	validatorPkg "gopkg.in/go-playground/validator.v8"
	"github.com/Confialink/wallet-pkg-types"
	"github.com/gin-gonic/gin"
	"github.com/Confialink/wallet-pkg-errors"
)

// request for validation
type Request struct {
	CategoryId int32
	Name       string
	CustomForm *custom_form.Form `json:"-"`
}

// implement custom_form.FormContainer interface
func (r Request) GetCustomForm() *custom_form.Form {
	return r.CustomForm
}

// Validator service
type Validator struct {
	validator        *validatorPkg.Validate
	formHelper       *custom_form.Helper
	validatorFactory *custom_form.ValidatorFactory
}

// You can register your own factories
// validatorFactory.Register(CustomFactory)
func NewValidator(validatorFactory *custom_form.ValidatorFactory) *Validator {
	vldtr := validatorPkg.New(&validatorPkg.Config{TagName: "validate"})
	formHelper := &custom_form.Helper{}
	
	return &Validator{vldtr, formHelper, validatorFactory}
}

// attach validator callback and validate struct
// `options` is some data which we can use into validators 
func (v *Validator) ValidateStruct(r *Request, options types.DataJSON) error {
	// make form
	form := &custom_form.Form{}
	
	// make field #1
	fieldCategory := custom_form.NewInt32Field("CategoryId")
	form.AddField(fieldCategory) // add field to the form
	
	// make field #2
	fieldName := custom_form.NewStringField("Name")
	condition := &custom_form.Condition{
		FieldName: fieldCategory.Name, // condition depends from this field
		Values:    []interface{}{12, 13}, // it means that validator will work only if r.Value is equal 12 or 13
		Type:      custom_form.ConditionTypeIn, // type of condition
	}
	// make conditions for validator
	conditions := make([]*custom_form.Condition, 1)
	conditions = append(conditions, condition)
	
	// make validator
	vldtr, err := v.validatorFactory.Make("max", conditions, options)
	if err != nil {
		return  err
	}
	fieldName.AddValidator(vldtr) // attach the validator to the field
	form.AddField(fieldName) // add the field to the form
    	
	// create validation callback. you can use your own
	customValidator := v.formHelper.MakeValidationCallback()
	
	// register callback in validator
	// it works only one time, after first call of validator.Struct the struct is cached in the v8.validator
	// so we cannot register a new callback every time
	// so we put custom form into the struct
	v.validator.RegisterStructValidation(customValidator, r)
	r.CustomForm = form
	// validate struct
	if err := v.validator.Struct(r); err != nil {
		return err
	}
	return nil
}

// action handler
func UpdateOne(c *gin.Context) {
	req := Request{}
	if err := c.ShouldBindJSON(&req); err != nil {
		errors.AddShouldBindError(c, err)
		return
	}

	validatorFactory := custom_form.NewValidatorFactory()
	// you can register your own validator factory with your validators
	// validatorFactory.Register(CustomValidatorFactory)
	validator := NewValidator(validatorFactory)
	options := types.DataJSON{}
	options.Set("value", 255) // it can be used in validators
	if err := validator.ValidateStruct(&req, options); err != nil {
		errors.AddShouldBindError(c, err)
		return
	}
}

```  
