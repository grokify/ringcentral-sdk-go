package definitions

type MeetingInfo struct {
	Topic                  string              `json:"topic,omitempty"`
	MeetingType            string              `json:"meetingType,omitempty"`
	Schedule               MeetingScheduleInfo `json:"schedule,omitempty"`
	StartParticipantsVideo bool                `json:"startParticipantsVideo,omitempty"`
	AudioOptions           `json:"audioOptions,omitempty"`
	StartHostVideo         bool      `json:"startHostVideo,omitempty"`
	Uri                    string    `json:"uri,omitempty"`
	Id                     string    `json:"id,omitempty"`
	Password               string    `json:"password,omitempty"`
	Status                 string    `json:"status,omitempty"`
	Links                  LinksInfo `json:"links,omitempty"`
	AllowJoinBeforeHost    bool      `json:"allowJoinBeforeHost,omitempty"`
}
