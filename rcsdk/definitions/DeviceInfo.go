package definitions

type DeviceInfo struct {
	PhoneLines              PhoneLinesInfo           `json:"phoneLines,omitempty"`
	Id                      string                   `json:"id,omitempty"`
	Uri                     string                   `json:"uri,omitempty"`
	Sku                     string                   `json:"sku,omitempty"`
	Serial                  string                   `json:"serial,omitempty"`
	Model                   ModelInfo                `json:"model,omitempty"`
	Extension               DeviceInfo_ExtensionInfo `json:"extension,omitempty"`
	EmergencyServiceAddress EmergencyAddressInfo     `json:"emergencyServiceAddress,omitempty"`
	Shipping                ShippingInfo             `json:"shipping,omitempty"`
	Type                    string                   `json:"type,omitempty"`
	Name                    string                   `json:"name,omitempty"`
	ComputerName            string                   `json:"computerName,omitempty"`
	BoxBillingId            int                      `json:"boxBillingId,omitempty"`
}
