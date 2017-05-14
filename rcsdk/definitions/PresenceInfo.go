package definitions

type PresenceInfo struct {
	Uri                 string                     `json:"uri,omitempty"`
	AllowSeeMyPresence  bool                       `json:"allowSeeMyPresence,omitempty"`
	DndStatus           string                     `json:"dndStatus,omitempty"`
	PresenceStatus      string                     `json:"presenceStatus,omitempty"`
	RingOnMonitoredCall bool                       `json:"ringOnMonitoredCall,omitempty"`
	Extension           PresenceInfo_ExtensionInfo `json:"extension,omitempty"`
	Message             string                     `json:"message,omitempty"`
	PickUpCallsOnHold   bool                       `json:"pickUpCallsOnHold,omitempty"`
	TelephonyStatus     string                     `json:"telephonyStatus,omitempty"`
	UserStatus          string                     `json:"userStatus,omitempty"`
}
