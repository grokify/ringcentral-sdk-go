package definitions

type CustomGreetingInfo struct {
	Uri         string `json:"uri,omitempty"`
	Id          string `json:"id,omitempty"`
	Type        string `json:"type,omitempty"`
	ContentType string `json:"contentType,omitempty"`
	ContentUri  string `json:"contentUri,omitempty"`
}
