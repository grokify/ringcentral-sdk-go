package definitions

type LocationInfo struct {
	State    string `json:"state,omitempty"`
	Uri      string `json:"uri,omitempty"`
	AreaCode string `json:"areaCode,omitempty"`
	City     string `json:"city,omitempty"`
	Npa      string `json:"npa,omitempty"`
	Nxx      string `json:"nxx,omitempty"`
}
