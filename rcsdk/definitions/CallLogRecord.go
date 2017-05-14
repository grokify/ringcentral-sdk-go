package definitions

type CallLogRecord struct {
	Id               string `json:"id,omitempty"`
	SessionId        string `json:"sessionId,omitempty"`
	Transport        string `json:"transport,omitempty"`
	Action           string `json:"action,omitempty"`
	StartTime        string `json:"startTime,omitempty"`
	Duration         int    `json:"duration,omitempty"`
	Legs             `json:"legs,omitempty"`
	From             CallerInfo    `json:"from,omitempty"`
	Type             string        `json:"type,omitempty"`
	LastModifiedTime string        `json:"lastModifiedTime,omitempty"`
	Uri              string        `json:"uri,omitempty"`
	To               CallerInfo    `json:"to,omitempty"`
	Direction        string        `json:"direction,omitempty"`
	Result           string        `json:"result,omitempty"`
	Recording        RecordingInfo `json:"recording,omitempty"`
}
