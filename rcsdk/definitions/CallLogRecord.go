package definitions

type CallLogRecord struct {
	Action           string        `json:"action,omitempty"`
	Duration         int           `json:"duration,omitempty"`
	LastModifiedTime string        `json:"lastModifiedTime,omitempty"`
	To               CallerInfo    `json:"to,omitempty"`
	Result           string        `json:"result,omitempty"`
	Recording        RecordingInfo `json:"recording,omitempty"`
	Legs             `json:"legs,omitempty"`
	Type             string     `json:"type,omitempty"`
	From             CallerInfo `json:"from,omitempty"`
	Direction        string     `json:"direction,omitempty"`
	StartTime        string     `json:"startTime,omitempty"`
	Transport        string     `json:"transport,omitempty"`
	SessionId        string     `json:"sessionId,omitempty"`
	Uri              string     `json:"uri,omitempty"`
	Id               string     `json:"id,omitempty"`
}
