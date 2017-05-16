package definitions

type ParsePhoneNumber_CountryInfo struct {
	CallingCode      string `json:"callingCode,omitempty"`
	EmergencyCalling bool   `json:"emergencyCalling,omitempty"`
	Id               string `json:"id,omitempty"`
	IsoCode          string `json:"isoCode,omitempty"`
	Name             string `json:"name,omitempty"`
	Uri              string `json:"uri,omitempty"`
}
