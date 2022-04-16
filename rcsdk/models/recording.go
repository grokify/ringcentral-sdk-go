package models

type RecordingInfo struct {
	URI         string `json:"uri"`
	ID          int64  `json:"id"`
	ContentType string `json:"contentType"`
	Duration    int    `json:"duration"`
}
