package definitions

type ParsePhoneNumber_CountryInfo struct {
	IsoCode          string `json:"isoCode,omitempty"`
	Name             string `json:"name,omitempty"`
	Id               string `json:"id,omitempty"`
	Uri              string `json:"uri,omitempty"`
	CallingCode      string `json:"callingCode,omitempty"`
	EmergencyCalling bool   `json:"emergencyCalling,omitempty"`
}
