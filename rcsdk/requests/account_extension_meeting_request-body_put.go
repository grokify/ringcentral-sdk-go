package requests

import (
	"github.com/grokify/ringcentral-sdk-go/rcsdk/definitions"
)

type AccountExtensionMeetingPutRequestBody struct {
	AllowJoinBeforeHost    bool                            `json:"allowJoinBeforeHost,omitempty"`
	AudioOptions           []string                        `json:"audioOptions,omitempty"`
	MeetingType            string                          `json:"meetingType,omitempty"`
	Password               string                          `json:"password,omitempty"`
	Schedule               definitions.MeetingScheduleInfo `json:"schedule,omitempty"`
	StartHostVideo         bool                            `json:"startHostVideo,omitempty"`
	StartParticipantsVideo bool                            `json:"startParticipantsVideo,omitempty"`
	Topic                  string                          `json:"topic,omitempty"`
}
