package definitions

type InstantMessageEvent struct {
	ReadStatus       string `json:"readStatus,omitempty"`
	Availability     string `json:"availability,omitempty"`
	MessageStatus    string `json:"messageStatus,omitempty"`
	To               `json:"to,omitempty"`
	From             InstantMessageEvent_CallerInfo `json:"from,omitempty"`
	CreationTime     string                         `json:"creationTime,omitempty"`
	LastModifiedTime string                         `json:"lastModifiedTime,omitempty"`
	Attachments      `json:"attachments,omitempty"`
	Subject          string `json:"subject,omitempty"`
	Direction        string `json:"direction,omitempty"`
	Id               string `json:"id,omitempty"`
	Type             string `json:"type,omitempty"`
	Priority         string `json:"priority,omitempty"`
	ConversationId   string `json:"conversationId,omitempty"`
}
