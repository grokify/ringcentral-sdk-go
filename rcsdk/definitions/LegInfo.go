package definitions

type LegInfo struct {
	Result    string                `json:"result,omitempty"`
	Recording RecordingInfo         `json:"recording,omitempty"`
	Duration  int                   `json:"duration,omitempty"`
	Extension LegInfo_ExtensionInfo `json:"extension,omitempty"`
	LegType   string                `json:"legType,omitempty"`
	StartTime string                `json:"startTime,omitempty"`
	To        CallerInfo            `json:"to,omitempty"`
	Transport string                `json:"transport,omitempty"`
	Action    string                `json:"action,omitempty"`
	Direction string                `json:"direction,omitempty"`
	Type      string                `json:"type,omitempty"`
	From      CallerInfo            `json:"from,omitempty"`
}
