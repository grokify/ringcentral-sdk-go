package definitions

type ExtensionListNotification struct {
	Body           ExtensionListEvent `json:"body,omitempty"`
	Event          string             `json:"event,omitempty"`
	SubscriptionID string             `json:"subscriptionId,omitempty"`
	Timestamp      string             `json:"timestamp,omitempty"`
	UURI           string             `json:"uuid,omitempty"`
}
