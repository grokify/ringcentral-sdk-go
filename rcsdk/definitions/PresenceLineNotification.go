package definitions

type PresenceLineNotification struct {
	Event          string            `json:"event,omitempty"`
	SubscriptionId string            `json:"subscriptionId,omitempty"`
	Timestamp      string            `json:"timestamp,omitempty"`
	Body           PresenceLineEvent `json:"body,omitempty"`
	Uuid           string            `json:"uuid,omitempty"`
}
