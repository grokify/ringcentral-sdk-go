package definitions

type PagingInfo struct {
	TotalPages    int `json:"totalPages,omitempty"`
	TotalElements int `json:"totalElements,omitempty"`
	Page          int `json:"page,omitempty"`
	PerPage       int `json:"perPage,omitempty"`
	PageStart     int `json:"pageStart,omitempty"`
	PageEnd       int `json:"pageEnd,omitempty"`
}
