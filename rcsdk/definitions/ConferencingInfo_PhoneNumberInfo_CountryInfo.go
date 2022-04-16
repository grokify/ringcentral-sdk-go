package definitions

type ConferencingInfo_PhoneNumberInfo_CountryInfo struct {
	CallingCode      string `json:"callingCode,omitempty"`
	EmergencyCalling bool   `json:"emergencyCalling,omitempty"`
	ID               string `json:"id,omitempty"`
	IsoCode          string `json:"isoCode,omitempty"`
	Name             string `json:"name,omitempty"`
	URI              string `json:"uri,omitempty"`
}
