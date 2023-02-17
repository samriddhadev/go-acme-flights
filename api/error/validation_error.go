package error

import "strings"

type ValidationField struct {
	FieldName string      `json:"field_name"`
	Message   string      `json:"message"`
	Details   interface{} `json:"details,omitempty"`
}

type ValidationError struct {
	Errors []ValidationField `json:"errors"`
}

func (e ValidationError) Error() string {
	var errString strings.Builder
	for _, validationfield := range e.Errors {
		errString.WriteString(validationfield.Message + "\n")
	}
	return errString.String()
}