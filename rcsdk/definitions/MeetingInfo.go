package definitions

type MeetingInfo struct {
	Password               string              `json:"password,omitempty"`
	Links                  LinksInfo           `json:"links,omitempty"`
	Schedule               MeetingScheduleInfo `json:"schedule,omitempty"`
	AllowJoinBeforeHost    bool                `json:"allowJoinBeforeHost,omitempty"`
	StartHostVideo         bool                `json:"startHostVideo,omitempty"`
	Uri                    string              `json:"uri,omitempty"`
	Topic                  string              `json:"topic,omitempty"`
	MeetingType            string              `json:"meetingType,omitempty"`
	StartParticipantsVideo bool                `json:"startParticipantsVideo,omitempty"`
	Id                     string              `json:"id,omitempty"`
	Status                 string              `json:"status,omitempty"`
	AudioOptions           `json:"audioOptions,omitempty"`
}
