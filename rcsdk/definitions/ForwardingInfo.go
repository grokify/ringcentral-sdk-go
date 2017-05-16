package definitions

type ForwardingInfo struct {
	NotifyAdminSoftPhones bool       `json:"notifyAdminSoftPhones,omitempty"`
	NotifyMySoftPhones    bool       `json:"notifyMySoftPhones,omitempty"`
	RingingMode           string     `json:"ringingMode,omitempty"`
	Rules                 []RuleInfo `json:"rules,omitempty"`
	SoftPhonesRingCount   int        `json:"softPhonesRingCount,omitempty"`
}
