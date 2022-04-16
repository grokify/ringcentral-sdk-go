package definitions

type CallLogRecord struct {
	Action           string        `json:"action,omitempty"`
	Direction        string        `json:"direction,omitempty"`
	Duration         int           `json:"duration,omitempty"`
	From             CallerInfo    `json:"from,omitempty"`
	ID               string        `json:"id,omitempty"`
	LastModifiedTime string        `json:"lastModifiedTime,omitempty"`
	Legs             []LegInfo     `json:"legs,omitempty"`
	Recording        RecordingInfo `json:"recording,omitempty"`
	Result           string        `json:"result,omitempty"`
	SessionID        string        `json:"sessionId,omitempty"`
	StartTime        string        `json:"startTime,omitempty"`
	To               CallerInfo    `json:"to,omitempty"`
	Transport        string        `json:"transport,omitempty"`
	Type             string        `json:"type,omitempty"`
	URI              string        `json:"uri,omitempty"`
}
