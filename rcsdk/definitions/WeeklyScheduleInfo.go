package definitions

type WeeklyScheduleInfo struct {
	Friday    []TimeInterval `json:"friday,omitempty"`
	Monday    []TimeInterval `json:"monday,omitempty"`
	Saturday  []TimeInterval `json:"saturday,omitempty"`
	Sunday    []TimeInterval `json:"sunday,omitempty"`
	Thursday  []TimeInterval `json:"thursday,omitempty"`
	Tuesday   []TimeInterval `json:"tuesday,omitempty"`
	Wednesday []TimeInterval `json:"wednesday,omitempty"`
}
