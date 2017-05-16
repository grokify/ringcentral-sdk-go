package definitions

type ServerInfo struct {
	ServerRevision string        `json:"serverRevision,omitempty"`
	Uri            string        `json:"uri,omitempty"`
	ApiVersions    []VersionInfo `json:"apiVersions,omitempty"`
	ServerVersion  string        `json:"serverVersion,omitempty"`
}
