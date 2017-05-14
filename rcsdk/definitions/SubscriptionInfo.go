package definitions

type SubscriptionInfo struct {
	EventFilters   `json:"eventFilters,omitempty"`
	ExpirationTime string       `json:"expirationTime,omitempty"`
	ExpiresIn      int          `json:"expiresIn,omitempty"`
	Status         string       `json:"status,omitempty"`
	CreationTime   string       `json:"creationTime,omitempty"`
	DeliveryMode   DeliveryMode `json:"deliveryMode,omitempty"`
	Id             string       `json:"id,omitempty"`
	Uri            string       `json:"uri,omitempty"`
}
