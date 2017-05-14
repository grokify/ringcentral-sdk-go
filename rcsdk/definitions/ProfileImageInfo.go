package definitions

type ProfileImageInfo struct {
	Etag         string `json:"etag,omitempty"`
	LastModified string `json:"lastModified,omitempty"`
	ContentType  string `json:"contentType,omitempty"`
	Scales       `json:"scales,omitempty"`
	Uri          string `json:"uri,omitempty"`
}
