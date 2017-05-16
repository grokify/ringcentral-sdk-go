package definitions

type MessageChange struct {
	Type         string `json:"type,omitempty"`
	NewCount     int    `json:"newCount,omitempty"`
	UpdatedCount int    `json:"updatedCount,omitempty"`
}
