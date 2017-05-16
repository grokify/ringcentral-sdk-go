package definitions

type DeviceInfo struct {
	BoxBillingId            int                      `json:"boxBillingId,omitempty"`
	ComputerName            string                   `json:"computerName,omitempty"`
	EmergencyServiceAddress EmergencyAddressInfo     `json:"emergencyServiceAddress,omitempty"`
	Extension               DeviceInfo_ExtensionInfo `json:"extension,omitempty"`
	Id                      string                   `json:"id,omitempty"`
	Model                   ModelInfo                `json:"model,omitempty"`
	Name                    string                   `json:"name,omitempty"`
	PhoneLines              PhoneLinesInfo           `json:"phoneLines,omitempty"`
	Serial                  string                   `json:"serial,omitempty"`
	Shipping                ShippingInfo             `json:"shipping,omitempty"`
	Sku                     string                   `json:"sku,omitempty"`
	Type                    string                   `json:"type,omitempty"`
	Uri                     string                   `json:"uri,omitempty"`
}
