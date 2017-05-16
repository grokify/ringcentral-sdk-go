package definitions

type MessageInfo struct {
	Attachments             []MessageAttachmentInfo  `json:"attachments,omitempty"`
	Availability            string                   `json:"availability,omitempty"`
	ConversationId          int                      `json:"conversationId,omitempty"`
	CreationTime            string                   `json:"creationTime,omitempty"`
	DeliveryErrorCode       string                   `json:"deliveryErrorCode,omitempty"`
	Direction               string                   `json:"direction,omitempty"`
	FaxPageCount            int                      `json:"faxPageCount,omitempty"`
	FaxResolution           string                   `json:"faxResolution,omitempty"`
	From                    MessageInfo_CallerInfo   `json:"from,omitempty"`
	Id                      string                   `json:"id,omitempty"`
	LastModifiedTime        string                   `json:"lastModifiedTime,omitempty"`
	MessageStatus           string                   `json:"messageStatus,omitempty"`
	PgToDepartment          bool                     `json:"pgToDepartment,omitempty"`
	Priority                string                   `json:"priority,omitempty"`
	ReadStatus              string                   `json:"readStatus,omitempty"`
	SmsDeliveryTime         string                   `json:"smsDeliveryTime,omitempty"`
	SmsSendingAttemptsCount int                      `json:"smsSendingAttemptsCount,omitempty"`
	Subject                 string                   `json:"subject,omitempty"`
	To                      []MessageInfo_CallerInfo `json:"to,omitempty"`
	Type                    string                   `json:"type,omitempty"`
	Uri                     string                   `json:"uri,omitempty"`
	VmTranscriptionStatus   string                   `json:"vmTranscriptionStatus,omitempty"`
}
