package definitions

type DetailedPresencewithSIPEvent_ActiveCallInfo struct {
	Direction       string  `json:"direction,omitempty"`
	From            string  `json:"from,omitempty"`
	Id              string  `json:"id,omitempty"`
	SessionId       string  `json:"sessionId,omitempty"`
	SipData         SIPData `json:"sipData,omitempty"`
	TelephonyStatus string  `json:"telephonyStatus,omitempty"`
	To              string  `json:"to,omitempty"`
}
