package definitions

type ExtensionListNotification struct {
	Body           ExtensionListEvent `json:"body,omitempty"`
	Event          string             `json:"event,omitempty"`
	SubscriptionId string             `json:"subscriptionId,omitempty"`
	Timestamp      string             `json:"timestamp,omitempty"`
	Uuid           string             `json:"uuid,omitempty"`
}
