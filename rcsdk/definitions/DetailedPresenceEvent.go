package definitions

type DetailedPresenceEvent struct {
	TerminationType     string `json:"terminationType,omitempty"`
	ActiveCalls         `json:"activeCalls,omitempty"`
	Sequence            int    `json:"sequence,omitempty"`
	DndStatus           string `json:"dndStatus,omitempty"`
	AllowSeeMyPresence  bool   `json:"allowSeeMyPresence,omitempty"`
	RingOnMonitoredCall bool   `json:"ringOnMonitoredCall,omitempty"`
	ExtensionId         string `json:"extensionId,omitempty"`
	TelephonyStatus     string `json:"telephonyStatus,omitempty"`
	PickUpCallsOnHold   bool   `json:"pickUpCallsOnHold,omitempty"`
	PresenceStatus      string `json:"presenceStatus,omitempty"`
	UserStatus          string `json:"userStatus,omitempty"`
}
