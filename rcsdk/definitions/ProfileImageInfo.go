package definitions

type ProfileImageInfo struct {
	Uri          string     `json:"uri,omitempty"`
	Etag         string     `json:"etag,omitempty"`
	LastModified string     `json:"lastModified,omitempty"`
	ContentType  string     `json:"contentType,omitempty"`
	Scales       []ImageUri `json:"scales,omitempty"`
}
