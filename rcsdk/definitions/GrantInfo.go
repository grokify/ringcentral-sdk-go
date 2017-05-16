package definitions

type GrantInfo struct {
	CallPickup     bool                    `json:"callPickup,omitempty"`
	CallMonitoring bool                    `json:"callMonitoring,omitempty"`
	Uri            string                  `json:"uri,omitempty"`
	Extension      GrantInfo_ExtensionInfo `json:"extension,omitempty"`
}
