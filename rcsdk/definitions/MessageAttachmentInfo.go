package definitions

type MessageAttachmentInfo struct {
	Type        string `json:"type,omitempty"`
	ContentType string `json:"contentType,omitempty"`
	VmDuration  int    `json:"vmDuration,omitempty"`
	Id          string `json:"id,omitempty"`
	Uri         string `json:"uri,omitempty"`
}
