package definitions

type InstantMessageEvent struct {
	Type             string `json:"type,omitempty"`
	Subject          string `json:"subject,omitempty"`
	Attachments      `json:"attachments,omitempty"`
	MessageStatus    string `json:"messageStatus,omitempty"`
	ReadStatus       string `json:"readStatus,omitempty"`
	Priority         string `json:"priority,omitempty"`
	CreationTime     string `json:"creationTime,omitempty"`
	Direction        string `json:"direction,omitempty"`
	Availability     string `json:"availability,omitempty"`
	ConversationId   string `json:"conversationId,omitempty"`
	Id               string `json:"id,omitempty"`
	To               `json:"to,omitempty"`
	From             InstantMessageEvent_CallerInfo `json:"from,omitempty"`
	LastModifiedTime string                         `json:"lastModifiedTime,omitempty"`
}
