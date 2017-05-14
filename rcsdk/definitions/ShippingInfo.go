package definitions

type ShippingInfo struct {
	Status         string `json:"status,omitempty"`
	Carrier        string `json:"carrier,omitempty"`
	TrackingNumber string `json:"trackingNumber,omitempty"`
	Method         `json:"method,omitempty"`
	Address        `json:"address,omitempty"`
}
