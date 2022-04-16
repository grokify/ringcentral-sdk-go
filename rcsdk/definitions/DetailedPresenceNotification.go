package definitions

type DetailedPresenceNotification struct {
	Body           DetailedPresenceEvent `json:"body,omitempty"`
	Event          string                `json:"event,omitempty"`
	SubscriptionId string                `json:"subscriptionId,omitempty"`
	Timestamp      string                `json:"timestamp,omitempty"`
	UUID           string                `json:"uuid,omitempty"`
}
