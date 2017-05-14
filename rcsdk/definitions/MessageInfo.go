package definitions

type MessageInfo struct {
	ConversationId          int    `json:"conversationId,omitempty"`
	CreationTime            string `json:"creationTime,omitempty"`
	DeliveryErrorCode       string `json:"deliveryErrorCode,omitempty"`
	PgToDepartment          bool   `json:"pgToDepartment,omitempty"`
	ReadStatus              string `json:"readStatus,omitempty"`
	Attachments             `json:"attachments,omitempty"`
	Availability            string                 `json:"availability,omitempty"`
	SmsDeliveryTime         string                 `json:"smsDeliveryTime,omitempty"`
	SmsSendingAttemptsCount int                    `json:"smsSendingAttemptsCount,omitempty"`
	From                    MessageInfo_CallerInfo `json:"from,omitempty"`
	LastModifiedTime        string                 `json:"lastModifiedTime,omitempty"`
	MessageStatus           string                 `json:"messageStatus,omitempty"`
	Priority                string                 `json:"priority,omitempty"`
	Subject                 string                 `json:"subject,omitempty"`
	VmTranscriptionStatus   string                 `json:"vmTranscriptionStatus,omitempty"`
	Id                      string                 `json:"id,omitempty"`
	Uri                     string                 `json:"uri,omitempty"`
	FaxResolution           string                 `json:"faxResolution,omitempty"`
	To                      `json:"to,omitempty"`
	Type                    string `json:"type,omitempty"`
	Direction               string `json:"direction,omitempty"`
	FaxPageCount            int    `json:"faxPageCount,omitempty"`
}
