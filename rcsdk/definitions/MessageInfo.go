package definitions

type MessageInfo struct {
	To                      `json:"to,omitempty"`
	DeliveryErrorCode       string `json:"deliveryErrorCode,omitempty"`
	MessageStatus           string `json:"messageStatus,omitempty"`
	Priority                string `json:"priority,omitempty"`
	Subject                 string `json:"subject,omitempty"`
	FaxResolution           string `json:"faxResolution,omitempty"`
	PgToDepartment          bool   `json:"pgToDepartment,omitempty"`
	VmTranscriptionStatus   string `json:"vmTranscriptionStatus,omitempty"`
	Id                      string `json:"id,omitempty"`
	Uri                     string `json:"uri,omitempty"`
	ConversationId          int    `json:"conversationId,omitempty"`
	FaxPageCount            int    `json:"faxPageCount,omitempty"`
	SmsDeliveryTime         string `json:"smsDeliveryTime,omitempty"`
	SmsSendingAttemptsCount int    `json:"smsSendingAttemptsCount,omitempty"`
	Type                    string `json:"type,omitempty"`
	CreationTime            string `json:"creationTime,omitempty"`
	Direction               string `json:"direction,omitempty"`
	LastModifiedTime        string `json:"lastModifiedTime,omitempty"`
	ReadStatus              string `json:"readStatus,omitempty"`
	Attachments             `json:"attachments,omitempty"`
	Availability            string                 `json:"availability,omitempty"`
	From                    MessageInfo_CallerInfo `json:"from,omitempty"`
}
