package definitions

type SyncInfo struct {
	SyncTime  string `json:"syncTime,omitempty"`
	SyncToken string `json:"syncToken,omitempty"`
	SyncType  string `json:"syncType,omitempty"`
}
