package definitions

type UserPermission struct {
	Permission UserPermissionInfo `json:"permission,omitempty"`
	Scopes     []string           `json:"scopes,omitempty"`
}
