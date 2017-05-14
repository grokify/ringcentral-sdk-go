package definitions

type StateInfo struct {
	Name    string                `json:"name,omitempty"`
	Id      string                `json:"id,omitempty"`
	Uri     string                `json:"uri,omitempty"`
	Country StateInfo_CountryInfo `json:"country,omitempty"`
	IsoCode string                `json:"isoCode,omitempty"`
}
