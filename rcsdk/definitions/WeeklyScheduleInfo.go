package definitions

type WeeklyScheduleInfo struct {
	Monday    []TimeInterval `json:"monday,omitempty"`
	Tuesday   []TimeInterval `json:"tuesday,omitempty"`
	Wednesday []TimeInterval `json:"wednesday,omitempty"`
	Thursday  []TimeInterval `json:"thursday,omitempty"`
	Friday    []TimeInterval `json:"friday,omitempty"`
	Saturday  []TimeInterval `json:"saturday,omitempty"`
	Sunday    []TimeInterval `json:"sunday,omitempty"`
}
