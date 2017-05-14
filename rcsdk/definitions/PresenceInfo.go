package definitions

type PresenceInfo struct {
	DndStatus           string                     `json:"dndStatus,omitempty"`
	Extension           PresenceInfo_ExtensionInfo `json:"extension,omitempty"`
	PickUpCallsOnHold   bool                       `json:"pickUpCallsOnHold,omitempty"`
	UserStatus          string                     `json:"userStatus,omitempty"`
	TelephonyStatus     string                     `json:"telephonyStatus,omitempty"`
	Uri                 string                     `json:"uri,omitempty"`
	AllowSeeMyPresence  bool                       `json:"allowSeeMyPresence,omitempty"`
	Message             string                     `json:"message,omitempty"`
	PresenceStatus      string                     `json:"presenceStatus,omitempty"`
	RingOnMonitoredCall bool                       `json:"ringOnMonitoredCall,omitempty"`
}
