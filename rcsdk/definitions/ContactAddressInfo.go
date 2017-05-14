package definitions

type ContactAddressInfo struct {
	Zip     string `json:"zip,omitempty"`
	Country string `json:"country,omitempty"`
	State   string `json:"state,omitempty"`
	City    string `json:"city,omitempty"`
	Street  string `json:"street,omitempty"`
}
