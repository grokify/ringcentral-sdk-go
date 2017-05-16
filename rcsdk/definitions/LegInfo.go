package definitions

type LegInfo struct {
	LegType   string                `json:"legType,omitempty"`
	StartTime string                `json:"startTime,omitempty"`
	Type      string                `json:"type,omitempty"`
	From      CallerInfo            `json:"from,omitempty"`
	Action    string                `json:"action,omitempty"`
	Extension LegInfo_ExtensionInfo `json:"extension,omitempty"`
	Result    string                `json:"result,omitempty"`
	To        CallerInfo            `json:"to,omitempty"`
	Transport string                `json:"transport,omitempty"`
	Recording RecordingInfo         `json:"recording,omitempty"`
	Direction string                `json:"direction,omitempty"`
	Duration  int                   `json:"duration,omitempty"`
}
