package definitions

type MeetingInfo struct {
	Password               string              `json:"password,omitempty"`
	AllowJoinBeforeHost    bool                `json:"allowJoinBeforeHost,omitempty"`
	StartParticipantsVideo bool                `json:"startParticipantsVideo,omitempty"`
	AudioOptions           []string            `json:"audioOptions,omitempty"`
	Uri                    string              `json:"uri,omitempty"`
	Id                     string              `json:"id,omitempty"`
	Topic                  string              `json:"topic,omitempty"`
	MeetingType            string              `json:"meetingType,omitempty"`
	Status                 string              `json:"status,omitempty"`
	Links                  LinksInfo           `json:"links,omitempty"`
	Schedule               MeetingScheduleInfo `json:"schedule,omitempty"`
	StartHostVideo         bool                `json:"startHostVideo,omitempty"`
}
