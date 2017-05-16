package definitions

type InstantMessageEvent struct {
	Attachments      []InstantMessageAttachmentInfo   `json:"attachments,omitempty"`
	Availability     string                           `json:"availability,omitempty"`
	ConversationId   string                           `json:"conversationId,omitempty"`
	CreationTime     string                           `json:"creationTime,omitempty"`
	Direction        string                           `json:"direction,omitempty"`
	From             InstantMessageEvent_CallerInfo   `json:"from,omitempty"`
	Id               string                           `json:"id,omitempty"`
	LastModifiedTime string                           `json:"lastModifiedTime,omitempty"`
	MessageStatus    string                           `json:"messageStatus,omitempty"`
	Priority         string                           `json:"priority,omitempty"`
	ReadStatus       string                           `json:"readStatus,omitempty"`
	Subject          string                           `json:"subject,omitempty"`
	To               []InstantMessageEvent_CallerInfo `json:"to,omitempty"`
	Type             string                           `json:"type,omitempty"`
}
