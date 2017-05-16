package definitions

type DetailedPresencewithSIPEvent struct {
	TelephonyStatus     string                                        `json:"telephonyStatus,omitempty"`
	TerminationType     string                                        `json:"terminationType,omitempty"`
	ActiveCalls         []DetailedPresencewithSIPEvent_ActiveCallInfo `json:"activeCalls,omitempty"`
	Sequence            int                                           `json:"sequence,omitempty"`
	PresenceStatus      string                                        `json:"presenceStatus,omitempty"`
	RingOnMonitoredCall bool                                          `json:"ringOnMonitoredCall,omitempty"`
	PickUpCallsOnHold   bool                                          `json:"pickUpCallsOnHold,omitempty"`
	ExtensionId         string                                        `json:"extensionId,omitempty"`
	DndStatus           string                                        `json:"dndStatus,omitempty"`
	AllowSeeMyPresence  bool                                          `json:"allowSeeMyPresence,omitempty"`
	UserStatus          string                                        `json:"userStatus,omitempty"`
}
