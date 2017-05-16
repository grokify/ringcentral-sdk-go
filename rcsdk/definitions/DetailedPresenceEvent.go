package definitions

type DetailedPresenceEvent struct {
	UserStatus          string                                 `json:"userStatus,omitempty"`
	AllowSeeMyPresence  bool                                   `json:"allowSeeMyPresence,omitempty"`
	RingOnMonitoredCall bool                                   `json:"ringOnMonitoredCall,omitempty"`
	PickUpCallsOnHold   bool                                   `json:"pickUpCallsOnHold,omitempty"`
	TerminationType     string                                 `json:"terminationType,omitempty"`
	ActiveCalls         []DetailedPresenceEvent_ActiveCallInfo `json:"activeCalls,omitempty"`
	Sequence            int                                    `json:"sequence,omitempty"`
	PresenceStatus      string                                 `json:"presenceStatus,omitempty"`
	DndStatus           string                                 `json:"dndStatus,omitempty"`
	ExtensionId         string                                 `json:"extensionId,omitempty"`
	TelephonyStatus     string                                 `json:"telephonyStatus,omitempty"`
}
