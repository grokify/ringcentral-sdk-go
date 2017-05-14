package definitions

type DetailedPresenceNotification struct {
	Uuid           string                `json:"uuid,omitempty"`
	Event          string                `json:"event,omitempty"`
	SubscriptionId string                `json:"subscriptionId,omitempty"`
	Timestamp      string                `json:"timestamp,omitempty"`
	Body           DetailedPresenceEvent `json:"body,omitempty"`
}
