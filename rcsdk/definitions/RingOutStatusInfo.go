package definitions

type RingOutStatusInfo struct {
	CallerStatus string `json:"callerStatus,omitempty"`
	CalleeStatus string `json:"calleeStatus,omitempty"`
	CallStatus   string `json:"callStatus,omitempty"`
}
