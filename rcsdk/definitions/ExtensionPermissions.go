package definitions

type ExtensionPermissions struct {
	Admin                PermissionInfo `json:"admin,omitempty"`
	InternationalCalling PermissionInfo `json:"internationalCalling,omitempty"`
}
