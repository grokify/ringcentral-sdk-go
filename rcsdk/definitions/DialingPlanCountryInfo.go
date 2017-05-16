package definitions

type DialingPlanCountryInfo struct {
	CallingCode string `json:"callingCode,omitempty"`
	Id          string `json:"id,omitempty"`
	IsoCode     string `json:"isoCode,omitempty"`
	Name        string `json:"name,omitempty"`
	Uri         string `json:"uri,omitempty"`
}
