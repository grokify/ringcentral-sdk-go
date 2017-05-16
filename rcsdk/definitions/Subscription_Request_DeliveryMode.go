package definitions

type Subscription_Request_DeliveryMode struct {
	Encryption    bool   `json:"encryption,omitempty"`
	TransportType string `json:"transportType,omitempty"`
}
