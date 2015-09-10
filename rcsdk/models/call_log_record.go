package models

type CallLogRecordsSimpleResponse struct {
	Uri        string                `json:"uri"`
	Records    []CallLogRecordSimple `json:"records"`
	Paging     Paging                `json:"paging"`
	Navigation Navigation            `json:"navigation"`
}

type CallLogRecordSimple struct {
	Uri       string    `json:"uri"`
	Id        string    `json:"id"`
	SessionId string    `json:"sessionId"`
	StartTime string    `json:"startTime"`
	Duration  int       `json:"duration"`
	Type      string    `json:"type"`
	Direction string    `json:"direction"`
	Action    string    `json:"action"`
	Result    string    `json:"result"`
	To        Contact   `json:"to"`
	From      Contact   `json:"from"`
	Recording Recording `json:"recording"`
}

type CallLogRecordDetailed struct{}
