package definitions

type ProfileImageInfo struct {
	ContentType  string `json:"contentType,omitempty"`
	Scales       `json:"scales,omitempty"`
	Uri          string `json:"uri,omitempty"`
	Etag         string `json:"etag,omitempty"`
	LastModified string `json:"lastModified,omitempty"`
}
