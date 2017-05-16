package definitions

type SubscriptionInfo struct {
	CreationTime   string       `json:"creationTime,omitempty"`
	DeliveryMode   DeliveryMode `json:"deliveryMode,omitempty"`
	EventFilters   []string     `json:"eventFilters,omitempty"`
	ExpirationTime string       `json:"expirationTime,omitempty"`
	ExpiresIn      int          `json:"expiresIn,omitempty"`
	Id             string       `json:"id,omitempty"`
	Status         string       `json:"status,omitempty"`
	Uri            string       `json:"uri,omitempty"`
}
