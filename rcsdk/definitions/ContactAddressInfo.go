package definitions

type ContactAddressInfo struct {
	State   string `json:"state,omitempty"`
	City    string `json:"city,omitempty"`
	Street  string `json:"street,omitempty"`
	Zip     string `json:"zip,omitempty"`
	Country string `json:"country,omitempty"`
}
