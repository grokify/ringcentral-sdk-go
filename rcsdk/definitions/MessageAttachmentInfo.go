package definitions

type MessageAttachmentInfo struct {
	ContentType string `json:"contentType,omitempty"`
	Id          string `json:"id,omitempty"`
	Type        string `json:"type,omitempty"`
	Uri         string `json:"uri,omitempty"`
	VmDuration  int    `json:"vmDuration,omitempty"`
}
