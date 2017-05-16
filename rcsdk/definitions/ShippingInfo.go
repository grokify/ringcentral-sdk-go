package definitions

type ShippingInfo struct {
	Address        []ShippingAddress `json:"address,omitempty"`
	Carrier        string            `json:"carrier,omitempty"`
	Method         []ShippingMethod  `json:"method,omitempty"`
	Status         string            `json:"status,omitempty"`
	TrackingNumber string            `json:"trackingNumber,omitempty"`
}
