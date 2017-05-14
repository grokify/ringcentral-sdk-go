package definitions

type CallLogInfo struct {
	Uri       string        `json:"uri,omitempty"`
	From      CallerInfo    `json:"from,omitempty"`
	To        CallerInfo    `json:"to,omitempty"`
	Result    string        `json:"result,omitempty"`
	Recording RecordingInfo `json:"recording,omitempty"`
	StartTime string        `json:"startTime,omitempty"`
	Duration  int           `json:"duration,omitempty"`
	Id        string        `json:"id,omitempty"`
	SessionId string        `json:"sessionId,omitempty"`
	Type      string        `json:"type,omitempty"`
	Direction string        `json:"direction,omitempty"`
	Action    string        `json:"action,omitempty"`
}
