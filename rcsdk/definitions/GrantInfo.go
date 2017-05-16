package definitions

type GrantInfo struct {
	CallMonitoring bool                    `json:"callMonitoring,omitempty"`
	CallPickup     bool                    `json:"callPickup,omitempty"`
	Extension      GrantInfo_ExtensionInfo `json:"extension,omitempty"`
	Uri            string                  `json:"uri,omitempty"`
}
