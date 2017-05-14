package definitions

type MessageInfo struct {
	FaxPageCount            int                    `json:"faxPageCount,omitempty"`
	From                    MessageInfo_CallerInfo `json:"from,omitempty"`
	SmsDeliveryTime         string                 `json:"smsDeliveryTime,omitempty"`
	To                      `json:"to,omitempty"`
	Type                    string `json:"type,omitempty"`
	Id                      string `json:"id,omitempty"`
	Availability            string `json:"availability,omitempty"`
	ReadStatus              string `json:"readStatus,omitempty"`
	Subject                 string `json:"subject,omitempty"`
	VmTranscriptionStatus   string `json:"vmTranscriptionStatus,omitempty"`
	Uri                     string `json:"uri,omitempty"`
	ConversationId          int    `json:"conversationId,omitempty"`
	Direction               string `json:"direction,omitempty"`
	PgToDepartment          bool   `json:"pgToDepartment,omitempty"`
	SmsSendingAttemptsCount int    `json:"smsSendingAttemptsCount,omitempty"`
	Attachments             `json:"attachments,omitempty"`
	CreationTime            string `json:"creationTime,omitempty"`
	DeliveryErrorCode       string `json:"deliveryErrorCode,omitempty"`
	FaxResolution           string `json:"faxResolution,omitempty"`
	LastModifiedTime        string `json:"lastModifiedTime,omitempty"`
	MessageStatus           string `json:"messageStatus,omitempty"`
	Priority                string `json:"priority,omitempty"`
}
