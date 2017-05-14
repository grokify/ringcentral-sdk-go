package definitions

type PresenceLineNotification struct {
	Uuid           string            `json:"uuid,omitempty"`
	Event          string            `json:"event,omitempty"`
	SubscriptionId string            `json:"subscriptionId,omitempty"`
	Timestamp      string            `json:"timestamp,omitempty"`
	Body           PresenceLineEvent `json:"body,omitempty"`
}
