package definitions

type AccountLimits struct {
	FreeSoftPhoneLinesPerExtension int `json:"freeSoftPhoneLinesPerExtension,omitempty"`
	MaxMonitoredExtensionsPerUser  int `json:"maxMonitoredExtensionsPerUser,omitempty"`
	MeetingSize                    int `json:"meetingSize,omitempty"`
}
