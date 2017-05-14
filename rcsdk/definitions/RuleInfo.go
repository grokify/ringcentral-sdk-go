package definitions

type RuleInfo struct {
	ForwardingNumbers `json:"forwardingNumbers,omitempty"`
	Index             int `json:"index,omitempty"`
	RingCount         int `json:"ringCount,omitempty"`
}
