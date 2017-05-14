package definitions

type RingOutStatusInfo struct {
	CallStatus   string `json:"callStatus,omitempty"`
	CallerStatus string `json:"callerStatus,omitempty"`
	CalleeStatus string `json:"calleeStatus,omitempty"`
}
