package definitions

type VersionInfo struct {
	ReleaseDate   string `json:"releaseDate,omitempty"`
	Uri           string `json:"uri,omitempty"`
	UriString     string `json:"uriString,omitempty"`
	VersionString string `json:"versionString,omitempty"`
}
