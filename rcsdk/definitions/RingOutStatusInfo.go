package definitions

type RingOutStatusInfo struct {
	CallStatus   string `json:"callStatus,omitempty"`
	CalleeStatus string `json:"calleeStatus,omitempty"`
	CallerStatus string `json:"callerStatus,omitempty"`
}
