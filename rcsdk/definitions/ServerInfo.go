package definitions

type ServerInfo struct {
	ServerVersion  string `json:"serverVersion,omitempty"`
	ServerRevision string `json:"serverRevision,omitempty"`
	Uri            string `json:"uri,omitempty"`
	ApiVersions    `json:"apiVersions,omitempty"`
}
