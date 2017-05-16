package definitions

type InstantMessageEvent struct {
	Availability     string                           `json:"availability,omitempty"`
	Id               string                           `json:"id,omitempty"`
	Attachments      []InstantMessageAttachmentInfo   `json:"attachments,omitempty"`
	ConversationId   string                           `json:"conversationId,omitempty"`
	LastModifiedTime string                           `json:"lastModifiedTime,omitempty"`
	Subject          string                           `json:"subject,omitempty"`
	ReadStatus       string                           `json:"readStatus,omitempty"`
	Priority         string                           `json:"priority,omitempty"`
	MessageStatus    string                           `json:"messageStatus,omitempty"`
	From             InstantMessageEvent_CallerInfo   `json:"from,omitempty"`
	CreationTime     string                           `json:"creationTime,omitempty"`
	Direction        string                           `json:"direction,omitempty"`
	To               []InstantMessageEvent_CallerInfo `json:"to,omitempty"`
	Type             string                           `json:"type,omitempty"`
}
