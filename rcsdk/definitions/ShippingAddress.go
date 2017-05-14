package definitions

type ShippingAddress struct {
	State        string `json:"state,omitempty"`
	Zip          string `json:"zip,omitempty"`
	Country      string `json:"country,omitempty"`
	CustomerName string `json:"customerName,omitempty"`
	Street       string `json:"street,omitempty"`
	Street2      string `json:"street2,omitempty"`
	City         string `json:"city,omitempty"`
}
