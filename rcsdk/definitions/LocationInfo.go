package definitions

type LocationInfo struct {
	City     string `json:"city,omitempty"`
	Npa      string `json:"npa,omitempty"`
	Nxx      string `json:"nxx,omitempty"`
	State    string `json:"state,omitempty"`
	Uri      string `json:"uri,omitempty"`
	AreaCode string `json:"areaCode,omitempty"`
}
