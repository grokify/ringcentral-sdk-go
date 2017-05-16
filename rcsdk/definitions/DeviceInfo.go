package definitions

type DeviceInfo struct {
	Sku                     string                   `json:"sku,omitempty"`
	Type                    string                   `json:"type,omitempty"`
	Name                    string                   `json:"name,omitempty"`
	ComputerName            string                   `json:"computerName,omitempty"`
	Extension               DeviceInfo_ExtensionInfo `json:"extension,omitempty"`
	EmergencyServiceAddress EmergencyAddressInfo     `json:"emergencyServiceAddress,omitempty"`
	BoxBillingId            int                      `json:"boxBillingId,omitempty"`
	Id                      string                   `json:"id,omitempty"`
	Uri                     string                   `json:"uri,omitempty"`
	Serial                  string                   `json:"serial,omitempty"`
	Model                   ModelInfo                `json:"model,omitempty"`
	PhoneLines              PhoneLinesInfo           `json:"phoneLines,omitempty"`
	Shipping                ShippingInfo             `json:"shipping,omitempty"`
}
