package definitions

type PresenceInfo struct {
	Uri                 string                     `json:"uri,omitempty"`
	AllowSeeMyPresence  bool                       `json:"allowSeeMyPresence,omitempty"`
	DndStatus           string                     `json:"dndStatus,omitempty"`
	Extension           PresenceInfo_ExtensionInfo `json:"extension,omitempty"`
	PickUpCallsOnHold   bool                       `json:"pickUpCallsOnHold,omitempty"`
	PresenceStatus      string                     `json:"presenceStatus,omitempty"`
	Message             string                     `json:"message,omitempty"`
	RingOnMonitoredCall bool                       `json:"ringOnMonitoredCall,omitempty"`
	TelephonyStatus     string                     `json:"telephonyStatus,omitempty"`
	UserStatus          string                     `json:"userStatus,omitempty"`
}
