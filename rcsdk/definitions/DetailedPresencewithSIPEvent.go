package definitions

type DetailedPresencewithSIPEvent struct {
	TerminationType     string `json:"terminationType,omitempty"`
	ActiveCalls         `json:"activeCalls,omitempty"`
	UserStatus          string `json:"userStatus,omitempty"`
	AllowSeeMyPresence  bool   `json:"allowSeeMyPresence,omitempty"`
	RingOnMonitoredCall bool   `json:"ringOnMonitoredCall,omitempty"`
	ExtensionId         string `json:"extensionId,omitempty"`
	TelephonyStatus     string `json:"telephonyStatus,omitempty"`
	Sequence            int    `json:"sequence,omitempty"`
	PresenceStatus      string `json:"presenceStatus,omitempty"`
	DndStatus           string `json:"dndStatus,omitempty"`
	PickUpCallsOnHold   bool   `json:"pickUpCallsOnHold,omitempty"`
}
