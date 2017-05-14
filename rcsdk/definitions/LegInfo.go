package definitions

type LegInfo struct {
	Action    string                `json:"action,omitempty"`
	Direction string                `json:"direction,omitempty"`
	Extension LegInfo_ExtensionInfo `json:"extension,omitempty"`
	Result    string                `json:"result,omitempty"`
	To        CallerInfo            `json:"to,omitempty"`
	Recording RecordingInfo         `json:"recording,omitempty"`
	Duration  int                   `json:"duration,omitempty"`
	LegType   string                `json:"legType,omitempty"`
	StartTime string                `json:"startTime,omitempty"`
	Type      string                `json:"type,omitempty"`
	From      CallerInfo            `json:"from,omitempty"`
	Transport string                `json:"transport,omitempty"`
}
