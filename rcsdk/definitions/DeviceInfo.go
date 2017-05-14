package definitions

type DeviceInfo struct {
	Shipping                ShippingInfo             `json:"shipping,omitempty"`
	BoxBillingId            int                      `json:"boxBillingId,omitempty"`
	Id                      string                   `json:"id,omitempty"`
	Serial                  string                   `json:"serial,omitempty"`
	ComputerName            string                   `json:"computerName,omitempty"`
	PhoneLines              PhoneLinesInfo           `json:"phoneLines,omitempty"`
	Model                   ModelInfo                `json:"model,omitempty"`
	Extension               DeviceInfo_ExtensionInfo `json:"extension,omitempty"`
	EmergencyServiceAddress EmergencyAddressInfo     `json:"emergencyServiceAddress,omitempty"`
	Uri                     string                   `json:"uri,omitempty"`
	Sku                     string                   `json:"sku,omitempty"`
	Type                    string                   `json:"type,omitempty"`
	Name                    string                   `json:"name,omitempty"`
}
