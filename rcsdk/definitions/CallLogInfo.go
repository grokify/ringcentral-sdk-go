package definitions

type CallLogInfo struct {
	Action    string        `json:"action,omitempty"`
	Direction string        `json:"direction,omitempty"`
	Duration  int           `json:"duration,omitempty"`
	From      CallerInfo    `json:"from,omitempty"`
	ID        string        `json:"id,omitempty"`
	Recording RecordingInfo `json:"recording,omitempty"`
	Result    string        `json:"result,omitempty"`
	SessionID string        `json:"sessionId,omitempty"`
	StartTime string        `json:"startTime,omitempty"`
	To        CallerInfo    `json:"to,omitempty"`
	Type      string        `json:"type,omitempty"`
	URI       string        `json:"uri,omitempty"`
}
