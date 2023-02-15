package error

type ValidationField struct {
	FieldName string `json:"field_name"`
	Message   string `json:"message"`
}

type APIError struct {
	Status int               `json:"status"`
	Errors []ValidationField `json:"errors"`
}
