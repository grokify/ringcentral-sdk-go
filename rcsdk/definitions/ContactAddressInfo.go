package definitions

type ContactAddressInfo struct {
	City    string `json:"city,omitempty"`
	Street  string `json:"street,omitempty"`
	Zip     string `json:"zip,omitempty"`
	Country string `json:"country,omitempty"`
	State   string `json:"state,omitempty"`
}
