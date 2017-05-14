package definitions

type UserPermission struct {
	Permission UserPermissionInfo `json:"permission,omitempty"`
	Scopes     `json:"scopes,omitempty"`
}
