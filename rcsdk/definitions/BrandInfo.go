package definitions

type BrandInfo struct {
	Id          string      `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	HomeCountry CountryInfo `json:"homeCountry,omitempty"`
}
