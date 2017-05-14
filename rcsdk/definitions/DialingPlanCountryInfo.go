package definitions

type DialingPlanCountryInfo struct {
	Id          string `json:"id,omitempty"`
	Uri         string `json:"uri,omitempty"`
	CallingCode string `json:"callingCode,omitempty"`
	IsoCode     string `json:"isoCode,omitempty"`
	Name        string `json:"name,omitempty"`
}
