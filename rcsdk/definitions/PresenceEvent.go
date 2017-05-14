package definitions

type PresenceEvent struct {
	RingOnMonitoredCall bool   `json:"ringOnMonitoredCall,omitempty"`
	PickUpCallsOnHold   bool   `json:"pickUpCallsOnHold,omitempty"`
	PresenceStatus      string `json:"presenceStatus,omitempty"`
	UserStatus          string `json:"userStatus,omitempty"`
	TerminationType     string `json:"terminationType,omitempty"`
	Sequence            int    `json:"sequence,omitempty"`
	DndStatus           string `json:"dndStatus,omitempty"`
	AllowSeeMyPresence  bool   `json:"allowSeeMyPresence,omitempty"`
	ExtensionId         string `json:"extensionId,omitempty"`
	TelephonyStatus     string `json:"telephonyStatus,omitempty"`
}
