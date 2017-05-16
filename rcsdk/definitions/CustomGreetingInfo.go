package definitions

type CustomGreetingInfo struct {
	ContentType string `json:"contentType,omitempty"`
	ContentUri  string `json:"contentUri,omitempty"`
	Id          string `json:"id,omitempty"`
	Type        string `json:"type,omitempty"`
	Uri         string `json:"uri,omitempty"`
}
