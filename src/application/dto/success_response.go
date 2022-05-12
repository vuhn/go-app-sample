package dto

type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

// NewSuccessResponse return new instance of SuccessResponse struct
func NewSuccessResponse(data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Success: true,
		Data:    data,
	}
}
