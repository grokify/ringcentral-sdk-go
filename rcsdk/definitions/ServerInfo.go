package definitions

type ServerInfo struct {
	ApiVersions    []VersionInfo `json:"apiVersions,omitempty"`
	ServerRevision string        `json:"serverRevision,omitempty"`
	ServerVersion  string        `json:"serverVersion,omitempty"`
	Uri            string        `json:"uri,omitempty"`
}
