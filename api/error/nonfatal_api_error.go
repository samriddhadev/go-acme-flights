package error

type NonFatalApiError struct {
	Message string `json:"Message"`
	Details string `json:"details,omitempty"`
}

func (e NonFatalApiError) Error() string {
	return e.Message
}