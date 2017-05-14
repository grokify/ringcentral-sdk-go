package definitions

type ExtensionListNotification struct {
	Uuid           string             `json:"uuid,omitempty"`
	Event          string             `json:"event,omitempty"`
	SubscriptionId string             `json:"subscriptionId,omitempty"`
	Timestamp      string             `json:"timestamp,omitempty"`
	Body           ExtensionListEvent `json:"body,omitempty"`
}
