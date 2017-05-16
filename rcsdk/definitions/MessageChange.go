package definitions

type MessageChange struct {
	NewCount     int    `json:"newCount,omitempty"`
	Type         string `json:"type,omitempty"`
	UpdatedCount int    `json:"updatedCount,omitempty"`
}
