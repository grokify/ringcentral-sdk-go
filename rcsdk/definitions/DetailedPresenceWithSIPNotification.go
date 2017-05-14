package definitions

type DetailedPresenceWithSIPNotification struct {
	Timestamp      string                       `json:"timestamp,omitempty"`
	Body           DetailedPresencewithSIPEvent `json:"body,omitempty"`
	Uuid           string                       `json:"uuid,omitempty"`
	Event          string                       `json:"event,omitempty"`
	SubscriptionId string                       `json:"subscriptionId,omitempty"`
}
