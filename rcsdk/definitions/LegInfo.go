package definitions

type LegInfo struct {
	Action    string                `json:"action,omitempty"`
	Direction string                `json:"direction,omitempty"`
	Duration  int                   `json:"duration,omitempty"`
	Extension LegInfo_ExtensionInfo `json:"extension,omitempty"`
	From      CallerInfo            `json:"from,omitempty"`
	LegType   string                `json:"legType,omitempty"`
	Recording RecordingInfo         `json:"recording,omitempty"`
	Result    string                `json:"result,omitempty"`
	StartTime string                `json:"startTime,omitempty"`
	To        CallerInfo            `json:"to,omitempty"`
	Transport string                `json:"transport,omitempty"`
	Type      string                `json:"type,omitempty"`
}
