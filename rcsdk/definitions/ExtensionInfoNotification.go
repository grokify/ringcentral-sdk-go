package definitions

type ExtensionInfoNotification struct {
	Uuid           string             `json:"uuid,omitempty"`
	Event          string             `json:"event,omitempty"`
	SubscriptionId string             `json:"subscriptionId,omitempty"`
	Timestamp      string             `json:"timestamp,omitempty"`
	Body           ExtensionInfoEvent `json:"body,omitempty"`
}
