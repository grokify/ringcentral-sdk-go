package definitions

type PresenceEvent struct {
	PickUpCallsOnHold   bool   `json:"pickUpCallsOnHold,omitempty"`
	ExtensionId         string `json:"extensionId,omitempty"`
	TelephonyStatus     string `json:"telephonyStatus,omitempty"`
	TerminationType     string `json:"terminationType,omitempty"`
	Sequence            int    `json:"sequence,omitempty"`
	UserStatus          string `json:"userStatus,omitempty"`
	AllowSeeMyPresence  bool   `json:"allowSeeMyPresence,omitempty"`
	PresenceStatus      string `json:"presenceStatus,omitempty"`
	DndStatus           string `json:"dndStatus,omitempty"`
	RingOnMonitoredCall bool   `json:"ringOnMonitoredCall,omitempty"`
}
