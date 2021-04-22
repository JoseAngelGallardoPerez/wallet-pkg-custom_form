package custom_form

// the struct which contains custom form
type FormContainer interface {
	// the struct, which we want validate, must contain custom form
	GetCustomForm() *Form
}

type Form struct {
	Fields []*Field
}

// add field to the form
func (f *Form) AddField(field *Field) {
	f.Fields = append(f.Fields, field)
}
