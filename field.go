package custom_form

const (
	FieldDataTypeString         = FieldDataType("string")
	FieldDataTypeInt32          = FieldDataType("int32")
	FieldDataTypeFloat32        = FieldDataType("float32")
	FieldDataTypeBool           = FieldDataType("bool")
	FieldDataTypeEmptyInterface = FieldDataType("")
)

type FieldDataType string

// a form contains fields
type Field struct {
	Name       string        `json:"name"`
	DataType   FieldDataType `json:"dataType"`
	Validators []Validator   `json:"validators"`
}

// create a new field
func NewField(name string, dataType FieldDataType) *Field {
	return &Field{Name: name, DataType: dataType}
}

// create a new string field
func NewStringField(name string) *Field {
	return &Field{Name: name, DataType: FieldDataTypeString}
}

// create a new in32 field
func NewInt32Field(name string) *Field {
	return &Field{Name: name, DataType: FieldDataTypeInt32}
}

// create a new float32 field
func NewFloat32Field(name string) *Field {
	return &Field{Name: name, DataType: FieldDataTypeFloat32}
}

// create a new bool field
func NewBoolField(name string) *Field {
	return &Field{Name: name, DataType: FieldDataTypeBool}
}

// create a new empty interface field
func NewEmptyInterfaceField(name string) *Field {
	return &Field{Name: name, DataType: FieldDataTypeEmptyInterface}
}

// add a new validator to the field
func (f *Field) AddValidator(validator Validator) {
	f.Validators = append(f.Validators, validator)
}

// set validators to the field
func (f *Field) SetValidators(validators []Validator) {
	f.Validators = validators
}
