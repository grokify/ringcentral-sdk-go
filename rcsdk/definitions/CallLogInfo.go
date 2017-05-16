package definitions

type CallLogInfo struct {
	Recording RecordingInfo `json:"recording,omitempty"`
	Uri       string        `json:"uri,omitempty"`
	SessionId string        `json:"sessionId,omitempty"`
	Direction string        `json:"direction,omitempty"`
	Duration  int           `json:"duration,omitempty"`
	Action    string        `json:"action,omitempty"`
	Result    string        `json:"result,omitempty"`
	StartTime string        `json:"startTime,omitempty"`
	Id        string        `json:"id,omitempty"`
	From      CallerInfo    `json:"from,omitempty"`
	To        CallerInfo    `json:"to,omitempty"`
	Type      string        `json:"type,omitempty"`
}
