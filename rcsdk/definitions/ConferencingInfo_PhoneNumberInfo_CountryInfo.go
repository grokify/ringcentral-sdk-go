package definitions

type ConferencingInfo_PhoneNumberInfo_CountryInfo struct {
	Id               string `json:"id,omitempty"`
	Uri              string `json:"uri,omitempty"`
	CallingCode      string `json:"callingCode,omitempty"`
	EmergencyCalling bool   `json:"emergencyCalling,omitempty"`
	IsoCode          string `json:"isoCode,omitempty"`
	Name             string `json:"name,omitempty"`
}
