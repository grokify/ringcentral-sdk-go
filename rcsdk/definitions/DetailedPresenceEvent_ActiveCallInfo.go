package definitions

type DetailedPresenceEvent_ActiveCallInfo struct {
	Direction       string `json:"direction,omitempty"`
	From            string `json:"from,omitempty"`
	Id              string `json:"id,omitempty"`
	SessionId       string `json:"sessionId,omitempty"`
	TelephonyStatus string `json:"telephonyStatus,omitempty"`
	TerminationType string `json:"terminationType,omitempty"`
	To              string `json:"to,omitempty"`
}
