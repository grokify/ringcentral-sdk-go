package definitions

type NavigationInfo struct {
	LastPage     Page `json:"lastPage,omitempty"`
	FirstPage    Page `json:"firstPage,omitempty"`
	NextPage     Page `json:"nextPage,omitempty"`
	PreviousPage Page `json:"previousPage,omitempty"`
}
