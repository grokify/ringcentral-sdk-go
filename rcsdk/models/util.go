package models

type Paging struct {
	Page      int `json:"page,omitempty"`
	PerPage   int `json:"perPage,omitempty"`
	PageStart int `json:"pageStart,omitempty"`
	PageEnd   int `json:"pageEnd,omitempty"`
}

type Navigation struct {
	FirstPage    URI `json:"firstPage,omitempty"`
	NextPage     URI `json:"nextPage,omitempty"`
	PreviousPage URI `json:"previousPage,omitempty"`
	LastPage     URI `json:"lastPage,omitempty"`
}

type URI struct {
	URI string `json:"uri"`
}
