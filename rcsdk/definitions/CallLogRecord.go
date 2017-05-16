package definitions

type CallLogRecord struct {
	SessionId        string        `json:"sessionId,omitempty"`
	Legs             []LegInfo     `json:"legs,omitempty"`
	Type             string        `json:"type,omitempty"`
	Duration         int           `json:"duration,omitempty"`
	Recording        RecordingInfo `json:"recording,omitempty"`
	Id               string        `json:"id,omitempty"`
	Uri              string        `json:"uri,omitempty"`
	To               CallerInfo    `json:"to,omitempty"`
	Transport        string        `json:"transport,omitempty"`
	Direction        string        `json:"direction,omitempty"`
	Result           string        `json:"result,omitempty"`
	LastModifiedTime string        `json:"lastModifiedTime,omitempty"`
	From             CallerInfo    `json:"from,omitempty"`
	Action           string        `json:"action,omitempty"`
	StartTime        string        `json:"startTime,omitempty"`
}
