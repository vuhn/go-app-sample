package dto

type PagingResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Paging  *Pagination `json:"paging"`
}

type Pagination struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// NewPagingResponse return new instance of PagingResponse struct
func NewPagingResponse(data interface{}, limit int, offset int, total int64) *PagingResponse {
	return &PagingResponse{
		Success: true,
		Data:    data,
		Paging: &Pagination{
			Limit:  limit,
			Offset: offset,
			Total:  total,
		},
	}
}
