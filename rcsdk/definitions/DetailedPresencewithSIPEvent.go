package definitions

type DetailedPresencewithSIPEvent struct {
	RingOnMonitoredCall bool   `json:"ringOnMonitoredCall,omitempty"`
	ExtensionId         string `json:"extensionId,omitempty"`
	TelephonyStatus     string `json:"telephonyStatus,omitempty"`
	TerminationType     string `json:"terminationType,omitempty"`
	ActiveCalls         `json:"activeCalls,omitempty"`
	PresenceStatus      string `json:"presenceStatus,omitempty"`
	UserStatus          string `json:"userStatus,omitempty"`
	AllowSeeMyPresence  bool   `json:"allowSeeMyPresence,omitempty"`
	Sequence            int    `json:"sequence,omitempty"`
	DndStatus           string `json:"dndStatus,omitempty"`
	PickUpCallsOnHold   bool   `json:"pickUpCallsOnHold,omitempty"`
}
