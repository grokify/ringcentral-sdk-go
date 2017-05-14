package definitions

type GrantInfo struct {
	CallMonitoring bool                    `json:"callMonitoring,omitempty"`
	Uri            string                  `json:"uri,omitempty"`
	Extension      GrantInfo_ExtensionInfo `json:"extension,omitempty"`
	CallPickup     bool                    `json:"callPickup,omitempty"`
}
