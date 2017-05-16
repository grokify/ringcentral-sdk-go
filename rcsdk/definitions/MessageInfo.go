package definitions

type MessageInfo struct {
	Availability            string                   `json:"availability,omitempty"`
	From                    MessageInfo_CallerInfo   `json:"from,omitempty"`
	PgToDepartment          bool                     `json:"pgToDepartment,omitempty"`
	SmsDeliveryTime         string                   `json:"smsDeliveryTime,omitempty"`
	Subject                 string                   `json:"subject,omitempty"`
	Attachments             []MessageAttachmentInfo  `json:"attachments,omitempty"`
	Uri                     string                   `json:"uri,omitempty"`
	ConversationId          int                      `json:"conversationId,omitempty"`
	Direction               string                   `json:"direction,omitempty"`
	FaxResolution           string                   `json:"faxResolution,omitempty"`
	Type                    string                   `json:"type,omitempty"`
	Id                      string                   `json:"id,omitempty"`
	Priority                string                   `json:"priority,omitempty"`
	ReadStatus              string                   `json:"readStatus,omitempty"`
	SmsSendingAttemptsCount int                      `json:"smsSendingAttemptsCount,omitempty"`
	To                      []MessageInfo_CallerInfo `json:"to,omitempty"`
	VmTranscriptionStatus   string                   `json:"vmTranscriptionStatus,omitempty"`
	DeliveryErrorCode       string                   `json:"deliveryErrorCode,omitempty"`
	FaxPageCount            int                      `json:"faxPageCount,omitempty"`
	LastModifiedTime        string                   `json:"lastModifiedTime,omitempty"`
	MessageStatus           string                   `json:"messageStatus,omitempty"`
	CreationTime            string                   `json:"creationTime,omitempty"`
}
