package definitions

type PresenceEvent struct {
	AllowSeeMyPresence  bool   `json:"allowSeeMyPresence,omitempty"`
	DndStatus           string `json:"dndStatus,omitempty"`
	ExtensionId         string `json:"extensionId,omitempty"`
	PickUpCallsOnHold   bool   `json:"pickUpCallsOnHold,omitempty"`
	PresenceStatus      string `json:"presenceStatus,omitempty"`
	RingOnMonitoredCall bool   `json:"ringOnMonitoredCall,omitempty"`
	Sequence            int    `json:"sequence,omitempty"`
	TelephonyStatus     string `json:"telephonyStatus,omitempty"`
	TerminationType     string `json:"terminationType,omitempty"`
	UserStatus          string `json:"userStatus,omitempty"`
}
