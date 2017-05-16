package definitions

type ShippingInfo struct {
	Carrier        string            `json:"carrier,omitempty"`
	TrackingNumber string            `json:"trackingNumber,omitempty"`
	Method         []ShippingMethod  `json:"method,omitempty"`
	Address        []ShippingAddress `json:"address,omitempty"`
	Status         string            `json:"status,omitempty"`
}
