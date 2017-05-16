package definitions

type RuleInfo struct {
	ForwardingNumbers []RuleInfo_ForwardingNumberInfo `json:"forwardingNumbers,omitempty"`
	Index             int                             `json:"index,omitempty"`
	RingCount         int                             `json:"ringCount,omitempty"`
}
