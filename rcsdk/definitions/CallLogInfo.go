package definitions

type CallLogInfo struct {
	Uri       string        `json:"uri,omitempty"`
	Type      string        `json:"type,omitempty"`
	Action    string        `json:"action,omitempty"`
	Result    string        `json:"result,omitempty"`
	StartTime string        `json:"startTime,omitempty"`
	Duration  int           `json:"duration,omitempty"`
	Id        string        `json:"id,omitempty"`
	From      CallerInfo    `json:"from,omitempty"`
	To        CallerInfo    `json:"to,omitempty"`
	Direction string        `json:"direction,omitempty"`
	Recording RecordingInfo `json:"recording,omitempty"`
	SessionId string        `json:"sessionId,omitempty"`
}
