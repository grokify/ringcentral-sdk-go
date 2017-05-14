package definitions

type RingOutInfo struct {
	Id     string            `json:"id,omitempty"`
	Status RingOutStatusInfo `json:"status,omitempty"`
}
