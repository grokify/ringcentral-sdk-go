package definitions

type DetailedPresencewithSIPEvent struct {
	ActiveCalls         `json:"activeCalls,omitempty"`
	Sequence            int    `json:"sequence,omitempty"`
	PresenceStatus      string `json:"presenceStatus,omitempty"`
	AllowSeeMyPresence  bool   `json:"allowSeeMyPresence,omitempty"`
	ExtensionId         string `json:"extensionId,omitempty"`
	TelephonyStatus     string `json:"telephonyStatus,omitempty"`
	TerminationType     string `json:"terminationType,omitempty"`
	PickUpCallsOnHold   bool   `json:"pickUpCallsOnHold,omitempty"`
	UserStatus          string `json:"userStatus,omitempty"`
	DndStatus           string `json:"dndStatus,omitempty"`
	RingOnMonitoredCall bool   `json:"ringOnMonitoredCall,omitempty"`
}
