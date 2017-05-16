package requests

import (
	"github.com/grokify/ringcentral-sdk-go/rcsdk/definitions"
)

type AccountExtensionMeetingPostRequestBody struct {
	MeetingType            string                          `json:"meetingType,omitempty"`
	Password               string                          `json:"password,omitempty"`
	Schedule               definitions.MeetingScheduleInfo `json:"schedule,omitempty"`
	AllowJoinBeforeHost    bool                            `json:"allowJoinBeforeHost,omitempty"`
	StartHostVideo         bool                            `json:"startHostVideo,omitempty"`
	StartParticipantsVideo bool                            `json:"startParticipantsVideo,omitempty"`
	AudioOptions           []string                        `json:"audioOptions,omitempty"`
	Topic                  string                          `json:"topic,omitempty"`
}
