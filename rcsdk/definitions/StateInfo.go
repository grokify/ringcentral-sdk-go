package definitions

type StateInfo struct {
	Id      string                `json:"id,omitempty"`
	Uri     string                `json:"uri,omitempty"`
	Country StateInfo_CountryInfo `json:"country,omitempty"`
	IsoCode string                `json:"isoCode,omitempty"`
	Name    string                `json:"name,omitempty"`
}
