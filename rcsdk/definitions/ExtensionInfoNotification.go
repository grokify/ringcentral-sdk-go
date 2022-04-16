package definitions

type ExtensionInfoNotification struct {
	Body           ExtensionInfoEvent `json:"body,omitempty"`
	Event          string             `json:"event,omitempty"`
	SubscriptionID string             `json:"subscriptionId,omitempty"`
	Timestamp      string             `json:"timestamp,omitempty"`
	UUID           string             `json:"uuid,omitempty"`
}
