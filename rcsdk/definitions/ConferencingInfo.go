package definitions

type ConferencingInfo struct {
	AllowJoinBeforeHost bool                               `json:"allowJoinBeforeHost,omitempty"`
	HostCode            string                             `json:"hostCode,omitempty"`
	Mode                string                             `json:"mode,omitempty"`
	ParticipantCode     string                             `json:"participantCode,omitempty"`
	PhoneNumber         string                             `json:"phoneNumber,omitempty"`
	PhoneNumbers        []ConferencingInfo_PhoneNumberInfo `json:"phoneNumbers,omitempty"`
	TapToJoinURI        string                             `json:"tapToJoinUri,omitempty"`
	URI                 string                             `json:"uri,omitempty"`
}
