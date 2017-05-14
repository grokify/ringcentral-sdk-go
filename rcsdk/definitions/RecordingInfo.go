package definitions

type RecordingInfo struct {
	Id         string `json:"id,omitempty"`
	Uri        string `json:"uri,omitempty"`
	Type       string `json:"type,omitempty"`
	ContentUri string `json:"contentUri,omitempty"`
}
