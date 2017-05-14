package definitions

type LocationInfo struct {
	Uri      string `json:"uri,omitempty"`
	AreaCode string `json:"areaCode,omitempty"`
	City     string `json:"city,omitempty"`
	Npa      string `json:"npa,omitempty"`
	Nxx      string `json:"nxx,omitempty"`
	State    string `json:"state,omitempty"`
}
