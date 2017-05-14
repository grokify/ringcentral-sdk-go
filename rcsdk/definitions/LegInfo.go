package definitions

type LegInfo struct {
	Direction string                `json:"direction,omitempty"`
	StartTime string                `json:"startTime,omitempty"`
	Recording RecordingInfo         `json:"recording,omitempty"`
	LegType   string                `json:"legType,omitempty"`
	Type      string                `json:"type,omitempty"`
	Result    string                `json:"result,omitempty"`
	From      CallerInfo            `json:"from,omitempty"`
	To        CallerInfo            `json:"to,omitempty"`
	Action    string                `json:"action,omitempty"`
	Duration  int                   `json:"duration,omitempty"`
	Extension LegInfo_ExtensionInfo `json:"extension,omitempty"`
	Transport string                `json:"transport,omitempty"`
}
