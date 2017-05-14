package definitions

type SyncInfo struct {
	SyncType  string `json:"syncType,omitempty"`
	SyncToken string `json:"syncToken,omitempty"`
	SyncTime  string `json:"syncTime,omitempty"`
}
