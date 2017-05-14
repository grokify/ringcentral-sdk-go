package definitions

type ScheduleInfo struct {
	WeeklyRanges WeeklyScheduleInfo `json:"weeklyRanges,omitempty"`
	Ranges       RangesInfo         `json:"ranges,omitempty"`
	Ref          string             `json:"ref,omitempty"`
}
