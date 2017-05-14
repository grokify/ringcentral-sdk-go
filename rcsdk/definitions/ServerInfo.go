package definitions

type ServerInfo struct {
	Uri            string `json:"uri,omitempty"`
	ApiVersions    `json:"apiVersions,omitempty"`
	ServerVersion  string `json:"serverVersion,omitempty"`
	ServerRevision string `json:"serverRevision,omitempty"`
}
