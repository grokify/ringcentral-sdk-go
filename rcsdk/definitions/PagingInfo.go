package definitions

type PagingInfo struct {
	Page          int `json:"page,omitempty"`
	PageEnd       int `json:"pageEnd,omitempty"`
	PageStart     int `json:"pageStart,omitempty"`
	PerPage       int `json:"perPage,omitempty"`
	TotalElements int `json:"totalElements,omitempty"`
	TotalPages    int `json:"totalPages,omitempty"`
}
