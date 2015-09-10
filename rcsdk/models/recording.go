package models

type RecordingInfo struct {
	Uri         string `json:"uri"`
	Id          int64  `json:"id"`
	ContentType string `json:"contentType"`
	Duration    int    `json:"duration"`
}
