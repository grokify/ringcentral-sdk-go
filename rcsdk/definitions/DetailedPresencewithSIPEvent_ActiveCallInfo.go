package definitions

type DetailedPresencewithSIPEvent_ActiveCallInfo struct {
	Direction       string  `json:"direction,omitempty"`
	From            string  `json:"from,omitempty"`
	ID              string  `json:"id,omitempty"`
	SessionID       string  `json:"sessionId,omitempty"`
	SipData         SIPData `json:"sipData,omitempty"`
	TelephonyStatus string  `json:"telephonyStatus,omitempty"`
	To              string  `json:"to,omitempty"`
}
