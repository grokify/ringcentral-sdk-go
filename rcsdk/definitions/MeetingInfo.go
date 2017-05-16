package definitions

type MeetingInfo struct {
	AllowJoinBeforeHost    bool                `json:"allowJoinBeforeHost,omitempty"`
	AudioOptions           []string            `json:"audioOptions,omitempty"`
	Id                     string              `json:"id,omitempty"`
	Links                  LinksInfo           `json:"links,omitempty"`
	MeetingType            string              `json:"meetingType,omitempty"`
	Password               string              `json:"password,omitempty"`
	Schedule               MeetingScheduleInfo `json:"schedule,omitempty"`
	StartHostVideo         bool                `json:"startHostVideo,omitempty"`
	StartParticipantsVideo bool                `json:"startParticipantsVideo,omitempty"`
	Status                 string              `json:"status,omitempty"`
	Topic                  string              `json:"topic,omitempty"`
	Uri                    string              `json:"uri,omitempty"`
}
