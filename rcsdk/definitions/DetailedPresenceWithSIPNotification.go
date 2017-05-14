package definitions

type DetailedPresenceWithSIPNotification struct {
	Event          string                       `json:"event,omitempty"`
	SubscriptionId string                       `json:"subscriptionId,omitempty"`
	Timestamp      string                       `json:"timestamp,omitempty"`
	Body           DetailedPresencewithSIPEvent `json:"body,omitempty"`
	Uuid           string                       `json:"uuid,omitempty"`
}
