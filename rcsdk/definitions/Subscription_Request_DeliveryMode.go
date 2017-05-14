package definitions

type Subscription_Request_DeliveryMode struct {
	TransportType string `json:"transportType,omitempty"`
	Encryption    bool   `json:"encryption,omitempty"`
}
