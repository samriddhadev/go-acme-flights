package error

type FatalApiError struct {
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func (e FatalApiError) Error() string {
	return e.Message
}