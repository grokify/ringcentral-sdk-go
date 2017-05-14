package definitions

type MeetingInfo struct {
	Id                     string              `json:"id,omitempty"`
	MeetingType            string              `json:"meetingType,omitempty"`
	Status                 string              `json:"status,omitempty"`
	Schedule               MeetingScheduleInfo `json:"schedule,omitempty"`
	StartParticipantsVideo bool                `json:"startParticipantsVideo,omitempty"`
	Uri                    string              `json:"uri,omitempty"`
	Password               string              `json:"password,omitempty"`
	Links                  LinksInfo           `json:"links,omitempty"`
	AllowJoinBeforeHost    bool                `json:"allowJoinBeforeHost,omitempty"`
	StartHostVideo         bool                `json:"startHostVideo,omitempty"`
	AudioOptions           `json:"audioOptions,omitempty"`
	Topic                  string `json:"topic,omitempty"`
}
