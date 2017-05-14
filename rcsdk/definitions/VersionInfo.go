package definitions

type VersionInfo struct {
	UriString     string `json:"uriString,omitempty"`
	Uri           string `json:"uri,omitempty"`
	VersionString string `json:"versionString,omitempty"`
	ReleaseDate   string `json:"releaseDate,omitempty"`
}
