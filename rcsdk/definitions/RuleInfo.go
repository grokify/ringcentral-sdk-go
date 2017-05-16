package definitions

type RuleInfo struct {
	Index             int                             `json:"index,omitempty"`
	RingCount         int                             `json:"ringCount,omitempty"`
	ForwardingNumbers []RuleInfo_ForwardingNumberInfo `json:"forwardingNumbers,omitempty"`
}
