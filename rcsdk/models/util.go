package models

type Paging struct {
	Page      int `json:"page,omitempty"`
	PerPage   int `json:"perPage,omitempty"`
	PageStart int `json:"pageStart,omitempty"`
	PageEnd   int `json:"pageEnd,omitempty"`
}

type Navigation struct {
	FirstPage    Uri `json:"firstPage,omitempty"`
	NextPage     Uri `json:"nextPage,omitempty"`
	PreviousPage Uri `json:"previousPage,omitempty"`
	LastPage     Uri `json:"lastPage,omitempty"`
}

type Uri struct {
	Uri string `json:"uri"`
}
