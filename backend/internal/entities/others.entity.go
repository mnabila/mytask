package entities

type ApiResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type ApiError struct {
	StatusCode int
	Message    string
}

func (e ApiError) Error() string {
	return e.Message
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
