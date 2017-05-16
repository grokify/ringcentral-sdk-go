package definitions

type DetailedPresenceEvent_ActiveCallInfo struct {
	To              string `json:"to,omitempty"`
	TelephonyStatus string `json:"telephonyStatus,omitempty"`
	SessionId       string `json:"sessionId,omitempty"`
	TerminationType string `json:"terminationType,omitempty"`
	Id              string `json:"id,omitempty"`
	Direction       string `json:"direction,omitempty"`
	From            string `json:"from,omitempty"`
}
