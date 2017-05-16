package definitions

type ScheduleInfo struct {
	Ranges       RangesInfo         `json:"ranges,omitempty"`
	Ref          string             `json:"ref,omitempty"`
	WeeklyRanges WeeklyScheduleInfo `json:"weeklyRanges,omitempty"`
}
