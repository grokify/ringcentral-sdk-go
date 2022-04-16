package definitions

type DetailedPresenceWithSIPNotification struct {
	Body           DetailedPresencewithSIPEvent `json:"body,omitempty"`
	Event          string                       `json:"event,omitempty"`
	SubscriptionID string                       `json:"subscriptionId,omitempty"`
	Timestamp      string                       `json:"timestamp,omitempty"`
	UUID           string                       `json:"uuid,omitempty"`
}
