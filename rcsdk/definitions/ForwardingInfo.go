package definitions

type ForwardingInfo struct {
	NotifyMySoftPhones    bool   `json:"notifyMySoftPhones,omitempty"`
	NotifyAdminSoftPhones bool   `json:"notifyAdminSoftPhones,omitempty"`
	SoftPhonesRingCount   int    `json:"softPhonesRingCount,omitempty"`
	RingingMode           string `json:"ringingMode,omitempty"`
	Rules                 `json:"rules,omitempty"`
}
