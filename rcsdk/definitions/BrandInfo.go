package definitions

type BrandInfo struct {
	HomeCountry CountryInfo `json:"homeCountry,omitempty"`
	ID          string      `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
}
