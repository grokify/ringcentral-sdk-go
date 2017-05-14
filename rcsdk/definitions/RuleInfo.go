package definitions

type RuleInfo struct {
	Index             int `json:"index,omitempty"`
	RingCount         int `json:"ringCount,omitempty"`
	ForwardingNumbers `json:"forwardingNumbers,omitempty"`
}
