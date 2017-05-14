package definitions

type CallLogInfo struct {
	StartTime string        `json:"startTime,omitempty"`
	Recording RecordingInfo `json:"recording,omitempty"`
	Id        string        `json:"id,omitempty"`
	Uri       string        `json:"uri,omitempty"`
	From      CallerInfo    `json:"from,omitempty"`
	To        CallerInfo    `json:"to,omitempty"`
	Type      string        `json:"type,omitempty"`
	Result    string        `json:"result,omitempty"`
	SessionId string        `json:"sessionId,omitempty"`
	Direction string        `json:"direction,omitempty"`
	Action    string        `json:"action,omitempty"`
	Duration  int           `json:"duration,omitempty"`
}
