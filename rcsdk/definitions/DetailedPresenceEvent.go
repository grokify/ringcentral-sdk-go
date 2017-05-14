package definitions

type DetailedPresenceEvent struct {
	PresenceStatus      string `json:"presenceStatus,omitempty"`
	AllowSeeMyPresence  bool   `json:"allowSeeMyPresence,omitempty"`
	PickUpCallsOnHold   bool   `json:"pickUpCallsOnHold,omitempty"`
	ExtensionId         string `json:"extensionId,omitempty"`
	TelephonyStatus     string `json:"telephonyStatus,omitempty"`
	TerminationType     string `json:"terminationType,omitempty"`
	ActiveCalls         `json:"activeCalls,omitempty"`
	Sequence            int    `json:"sequence,omitempty"`
	UserStatus          string `json:"userStatus,omitempty"`
	DndStatus           string `json:"dndStatus,omitempty"`
	RingOnMonitoredCall bool   `json:"ringOnMonitoredCall,omitempty"`
}
