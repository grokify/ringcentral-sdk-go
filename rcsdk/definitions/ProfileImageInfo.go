package definitions

type ProfileImageInfo struct {
	ContentType  string     `json:"contentType,omitempty"`
	Etag         string     `json:"etag,omitempty"`
	LastModified string     `json:"lastModified,omitempty"`
	Scales       []ImageUri `json:"scales,omitempty"`
	Uri          string     `json:"uri,omitempty"`
}
