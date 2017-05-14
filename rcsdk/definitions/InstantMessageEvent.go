package definitions

type InstantMessageEvent struct {
	To               `json:"to,omitempty"`
	Type             string                         `json:"type,omitempty"`
	Availability     string                         `json:"availability,omitempty"`
	Subject          string                         `json:"subject,omitempty"`
	MessageStatus    string                         `json:"messageStatus,omitempty"`
	Priority         string                         `json:"priority,omitempty"`
	ConversationId   string                         `json:"conversationId,omitempty"`
	Id               string                         `json:"id,omitempty"`
	From             InstantMessageEvent_CallerInfo `json:"from,omitempty"`
	CreationTime     string                         `json:"creationTime,omitempty"`
	LastModifiedTime string                         `json:"lastModifiedTime,omitempty"`
	ReadStatus       string                         `json:"readStatus,omitempty"`
	Attachments      `json:"attachments,omitempty"`
	Direction        string `json:"direction,omitempty"`
}
