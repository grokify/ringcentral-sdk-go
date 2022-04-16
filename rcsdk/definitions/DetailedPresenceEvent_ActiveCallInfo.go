package definitions

type DetailedPresenceEvent_ActiveCallInfo struct {
	Direction       string `json:"direction,omitempty"`
	From            string `json:"from,omitempty"`
	ID              string `json:"id,omitempty"`
	SessionID       string `json:"sessionId,omitempty"`
	TelephonyStatus string `json:"telephonyStatus,omitempty"`
	TerminationType string `json:"terminationType,omitempty"`
	To              string `json:"to,omitempty"`
}
