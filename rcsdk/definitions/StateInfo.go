package definitions

type StateInfo struct {
	Country StateInfo_CountryInfo `json:"country,omitempty"`
	Id      string                `json:"id,omitempty"`
	IsoCode string                `json:"isoCode,omitempty"`
	Name    string                `json:"name,omitempty"`
	Uri     string                `json:"uri,omitempty"`
}
