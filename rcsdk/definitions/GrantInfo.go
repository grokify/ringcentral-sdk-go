package definitions

type GrantInfo struct {
	Uri            string                  `json:"uri,omitempty"`
	Extension      GrantInfo_ExtensionInfo `json:"extension,omitempty"`
	CallPickup     bool                    `json:"callPickup,omitempty"`
	CallMonitoring bool                    `json:"callMonitoring,omitempty"`
}
