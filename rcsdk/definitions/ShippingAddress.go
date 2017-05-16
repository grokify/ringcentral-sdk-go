package definitions

type ShippingAddress struct {
	City         string `json:"city,omitempty"`
	Country      string `json:"country,omitempty"`
	CustomerName string `json:"customerName,omitempty"`
	State        string `json:"state,omitempty"`
	Street       string `json:"street,omitempty"`
	Street2      string `json:"street2,omitempty"`
	Zip          string `json:"zip,omitempty"`
}
