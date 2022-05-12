package dto

type ErrorResponse struct {
	Success bool        `json:"success"`
	Errors  interface{} `json:"errors"`
}

// NewErrorResponse return new instance of ErrorResponse
func NewErrorResponse(err interface{}) *ErrorResponse {
	return &ErrorResponse{
		Success: false,
		Errors:  err,
	}
}
