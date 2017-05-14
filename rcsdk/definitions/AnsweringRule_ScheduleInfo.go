package definitions

type AnsweringRule_ScheduleInfo struct {
	WeeklyRanges WeeklyScheduleInfo `json:"weeklyRanges,omitempty"`
	Ranges       RangesInfo         `json:"ranges,omitempty"`
}
