package definitions

type BrandInfo struct {
	HomeCountry CountryInfo `json:"homeCountry,omitempty"`
	Id          string      `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
}
