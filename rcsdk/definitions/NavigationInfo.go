package definitions

type NavigationInfo struct {
	FirstPage    Page `json:"firstPage,omitempty"`
	LastPage     Page `json:"lastPage,omitempty"`
	NextPage     Page `json:"nextPage,omitempty"`
	PreviousPage Page `json:"previousPage,omitempty"`
}
