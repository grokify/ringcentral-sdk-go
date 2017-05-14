package definitions

type NavigationInfo struct {
	PreviousPage Page `json:"previousPage,omitempty"`
	LastPage     Page `json:"lastPage,omitempty"`
	FirstPage    Page `json:"firstPage,omitempty"`
	NextPage     Page `json:"nextPage,omitempty"`
}
