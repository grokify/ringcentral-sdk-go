package definitions

type ConferencingInfo struct {
	PhoneNumbers        `json:"phoneNumbers,omitempty"`
	Uri                 string `json:"uri,omitempty"`
	AllowJoinBeforeHost bool   `json:"allowJoinBeforeHost,omitempty"`
	HostCode            string `json:"hostCode,omitempty"`
	Mode                string `json:"mode,omitempty"`
	ParticipantCode     string `json:"participantCode,omitempty"`
	PhoneNumber         string `json:"phoneNumber,omitempty"`
	TapToJoinUri        string `json:"tapToJoinUri,omitempty"`
}
