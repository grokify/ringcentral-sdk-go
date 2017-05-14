package definitions

type MessageChange struct {
	UpdatedCount int    `json:"updatedCount,omitempty"`
	Type         string `json:"type,omitempty"`
	NewCount     int    `json:"newCount,omitempty"`
}
