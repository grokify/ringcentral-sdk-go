package definitions

type ForwardingInfo struct {
	SoftPhonesRingCount   int    `json:"softPhonesRingCount,omitempty"`
	RingingMode           string `json:"ringingMode,omitempty"`
	Rules                 `json:"rules,omitempty"`
	NotifyMySoftPhones    bool `json:"notifyMySoftPhones,omitempty"`
	NotifyAdminSoftPhones bool `json:"notifyAdminSoftPhones,omitempty"`
}
