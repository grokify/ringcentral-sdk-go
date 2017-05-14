package definitions

type PagingInfo struct {
	TotalElements int `json:"totalElements,omitempty"`
	Page          int `json:"page,omitempty"`
	PerPage       int `json:"perPage,omitempty"`
	PageStart     int `json:"pageStart,omitempty"`
	PageEnd       int `json:"pageEnd,omitempty"`
	TotalPages    int `json:"totalPages,omitempty"`
}
