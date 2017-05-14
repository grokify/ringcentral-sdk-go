package definitions

type AccountLimits struct {
	FreeSoftPhoneLinesPerExtension int `json:"freeSoftPhoneLinesPerExtension,omitempty"`
	MeetingSize                    int `json:"meetingSize,omitempty"`
	MaxMonitoredExtensionsPerUser  int `json:"maxMonitoredExtensionsPerUser,omitempty"`
}
