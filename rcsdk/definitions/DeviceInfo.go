package definitions

type DeviceInfo struct {
	BoxBillingID            int                      `json:"boxBillingId,omitempty"`
	ComputerName            string                   `json:"computerName,omitempty"`
	EmergencyServiceAddress EmergencyAddressInfo     `json:"emergencyServiceAddress,omitempty"`
	Extension               DeviceInfo_ExtensionInfo `json:"extension,omitempty"`
	ID                      string                   `json:"id,omitempty"`
	Model                   ModelInfo                `json:"model,omitempty"`
	Name                    string                   `json:"name,omitempty"`
	PhoneLines              PhoneLinesInfo           `json:"phoneLines,omitempty"`
	Serial                  string                   `json:"serial,omitempty"`
	Shipping                ShippingInfo             `json:"shipping,omitempty"`
	Sku                     string                   `json:"sku,omitempty"`
	Type                    string                   `json:"type,omitempty"`
	URI                     string                   `json:"uri,omitempty"`
}
