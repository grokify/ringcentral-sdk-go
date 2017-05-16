package definitions

type VersionInfo struct {
	Uri           string `json:"uri,omitempty"`
	VersionString string `json:"versionString,omitempty"`
	ReleaseDate   string `json:"releaseDate,omitempty"`
	UriString     string `json:"uriString,omitempty"`
}
