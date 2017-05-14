package definitions

type BrandInfo struct {
	Name        string      `json:"name,omitempty"`
	HomeCountry CountryInfo `json:"homeCountry,omitempty"`
	Id          string      `json:"id,omitempty"`
}
