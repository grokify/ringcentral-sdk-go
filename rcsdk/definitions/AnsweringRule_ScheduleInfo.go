package definitions

type AnsweringRule_ScheduleInfo struct {
	Ranges       RangesInfo         `json:"ranges,omitempty"`
	WeeklyRanges WeeklyScheduleInfo `json:"weeklyRanges,omitempty"`
}
