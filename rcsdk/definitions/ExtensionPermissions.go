package definitions

type ExtensionPermissions struct {
	InternationalCalling PermissionInfo `json:"internationalCalling,omitempty"`
	Admin                PermissionInfo `json:"admin,omitempty"`
}
