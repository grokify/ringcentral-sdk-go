package definitions

type PresenceEvent struct {
	TerminationType     string `json:"terminationType,omitempty"`
	UserStatus          string `json:"userStatus,omitempty"`
	DndStatus           string `json:"dndStatus,omitempty"`
	RingOnMonitoredCall bool   `json:"ringOnMonitoredCall,omitempty"`
	PickUpCallsOnHold   bool   `json:"pickUpCallsOnHold,omitempty"`
	ExtensionId         string `json:"extensionId,omitempty"`
	TelephonyStatus     string `json:"telephonyStatus,omitempty"`
	Sequence            int    `json:"sequence,omitempty"`
	PresenceStatus      string `json:"presenceStatus,omitempty"`
	AllowSeeMyPresence  bool   `json:"allowSeeMyPresence,omitempty"`
}
