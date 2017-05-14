package definitions

type ExtensionInfoNotification struct {
	Timestamp      string             `json:"timestamp,omitempty"`
	Body           ExtensionInfoEvent `json:"body,omitempty"`
	Uuid           string             `json:"uuid,omitempty"`
	Event          string             `json:"event,omitempty"`
	SubscriptionId string             `json:"subscriptionId,omitempty"`
}
