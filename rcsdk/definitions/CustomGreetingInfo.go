package definitions

type CustomGreetingInfo struct {
	ContentType string `json:"contentType,omitempty"`
	ContentURI  string `json:"contentUri,omitempty"`
	ID          string `json:"id,omitempty"`
	Type        string `json:"type,omitempty"`
	URI         string `json:"uri,omitempty"`
}
