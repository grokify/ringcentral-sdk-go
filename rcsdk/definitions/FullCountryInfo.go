package definitions

type FullCountryInfo struct {
	CallingCode      string `json:"callingCode,omitempty"`
	EmergencyCalling bool   `json:"emergencyCalling,omitempty"`
	Id               string `json:"id,omitempty"`
	IsoCode          string `json:"isoCode,omitempty"`
	LoginAllowed     bool   `json:"loginAllowed,omitempty"`
	Name             string `json:"name,omitempty"`
	NumberSelling    bool   `json:"numberSelling,omitempty"`
	Uri              string `json:"uri,omitempty"`
}
