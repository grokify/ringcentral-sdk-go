package definitions

type WeeklyScheduleInfo struct {
	Sunday    `json:"sunday,omitempty"`
	Monday    `json:"monday,omitempty"`
	Tuesday   `json:"tuesday,omitempty"`
	Wednesday `json:"wednesday,omitempty"`
	Thursday  `json:"thursday,omitempty"`
	Friday    `json:"friday,omitempty"`
	Saturday  `json:"saturday,omitempty"`
}
