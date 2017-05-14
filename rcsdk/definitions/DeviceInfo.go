package definitions

type DeviceInfo struct {
	Type                    string                   `json:"type,omitempty"`
	Id                      string                   `json:"id,omitempty"`
	Uri                     string                   `json:"uri,omitempty"`
	Serial                  string                   `json:"serial,omitempty"`
	ComputerName            string                   `json:"computerName,omitempty"`
	Model                   ModelInfo                `json:"model,omitempty"`
	Extension               DeviceInfo_ExtensionInfo `json:"extension,omitempty"`
	EmergencyServiceAddress EmergencyAddressInfo     `json:"emergencyServiceAddress,omitempty"`
	PhoneLines              PhoneLinesInfo           `json:"phoneLines,omitempty"`
	Sku                     string                   `json:"sku,omitempty"`
	Name                    string                   `json:"name,omitempty"`
	Shipping                ShippingInfo             `json:"shipping,omitempty"`
	BoxBillingId            int                      `json:"boxBillingId,omitempty"`
}
