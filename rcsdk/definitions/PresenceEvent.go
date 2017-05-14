package definitions

type PresenceEvent struct {
	PresenceStatus      string `json:"presenceStatus,omitempty"`
	ExtensionId         string `json:"extensionId,omitempty"`
	TerminationType     string `json:"terminationType,omitempty"`
	UserStatus          string `json:"userStatus,omitempty"`
	DndStatus           string `json:"dndStatus,omitempty"`
	AllowSeeMyPresence  bool   `json:"allowSeeMyPresence,omitempty"`
	RingOnMonitoredCall bool   `json:"ringOnMonitoredCall,omitempty"`
	PickUpCallsOnHold   bool   `json:"pickUpCallsOnHold,omitempty"`
	TelephonyStatus     string `json:"telephonyStatus,omitempty"`
	Sequence            int    `json:"sequence,omitempty"`
}
