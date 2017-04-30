package models

type CallLogRecordsSimpleResponse struct {
	Uri        string                `json:"uri"`
	Records    []CallLogRecordSimple `json:"records"`
	Paging     Paging                `json:"paging"`
	Navigation Navigation            `json:"navigation"`
}

type CallLogRecordSimple struct {
	Uri       string        `json:"uri"`
	Id        string        `json:"id"`
	SessionId string        `json:"sessionId"`
	StartTime string        `json:"startTime"`
	Duration  int           `json:"duration"`
	Type      string        `json:"type"`
	Direction string        `json:"direction"`
	Action    string        `json:"action"`
	Result    string        `json:"result"`
	To        CallerInfo    `json:"to"`
	From      CallerInfo    `json:"from"`
	Recording RecordingInfo `json:"recording"`
}

type CallLogRecordDetailed struct{}
