package definitions

type PresenceInfo struct {
	AllowSeeMyPresence  bool                       `json:"allowSeeMyPresence,omitempty"`
	DndStatus           string                     `json:"dndStatus,omitempty"`
	Extension           PresenceInfo_ExtensionInfo `json:"extension,omitempty"`
	Message             string                     `json:"message,omitempty"`
	PickUpCallsOnHold   bool                       `json:"pickUpCallsOnHold,omitempty"`
	PresenceStatus      string                     `json:"presenceStatus,omitempty"`
	RingOnMonitoredCall bool                       `json:"ringOnMonitoredCall,omitempty"`
	TelephonyStatus     string                     `json:"telephonyStatus,omitempty"`
	Uri                 string                     `json:"uri,omitempty"`
	UserStatus          string                     `json:"userStatus,omitempty"`
}
