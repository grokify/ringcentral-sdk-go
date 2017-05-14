package definitions

type CallLogRecord struct {
	Id               string `json:"id,omitempty"`
	LastModifiedTime string `json:"lastModifiedTime,omitempty"`
	SessionId        string `json:"sessionId,omitempty"`
	Type             string `json:"type,omitempty"`
	Action           string `json:"action,omitempty"`
	Result           string `json:"result,omitempty"`
	Legs             `json:"legs,omitempty"`
	StartTime        string        `json:"startTime,omitempty"`
	Duration         int           `json:"duration,omitempty"`
	Recording        RecordingInfo `json:"recording,omitempty"`
	Uri              string        `json:"uri,omitempty"`
	From             CallerInfo    `json:"from,omitempty"`
	To               CallerInfo    `json:"to,omitempty"`
	Direction        string        `json:"direction,omitempty"`
	Transport        string        `json:"transport,omitempty"`
}
