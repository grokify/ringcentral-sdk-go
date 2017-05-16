package definitions

type CustomGreetingInfo struct {
	Id          string `json:"id,omitempty"`
	Type        string `json:"type,omitempty"`
	ContentType string `json:"contentType,omitempty"`
	ContentUri  string `json:"contentUri,omitempty"`
	Uri         string `json:"uri,omitempty"`
}
