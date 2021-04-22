package custom_form

const ConditionTypeIn = ConditionType("in")

type ConditionType string

// We can validate field in a form by some conditions.
type Condition struct {
	FieldName string        `json:"fieldName"` // validator depends from this field in the form
	Values    []interface{} `json:"values"`    // values of FieldName in the form
	Type      ConditionType `json:"type"`      // type of condition.
}
