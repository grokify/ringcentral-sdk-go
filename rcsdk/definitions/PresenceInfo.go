package definitions

type PresenceInfo struct {
	Uri                 string                     `json:"uri,omitempty"`
	AllowSeeMyPresence  bool                       `json:"allowSeeMyPresence,omitempty"`
	DndStatus           string                     `json:"dndStatus,omitempty"`
	Extension           PresenceInfo_ExtensionInfo `json:"extension,omitempty"`
	TelephonyStatus     string                     `json:"telephonyStatus,omitempty"`
	Message             string                     `json:"message,omitempty"`
	PickUpCallsOnHold   bool                       `json:"pickUpCallsOnHold,omitempty"`
	PresenceStatus      string                     `json:"presenceStatus,omitempty"`
	RingOnMonitoredCall bool                       `json:"ringOnMonitoredCall,omitempty"`
	UserStatus          string                     `json:"userStatus,omitempty"`
}
