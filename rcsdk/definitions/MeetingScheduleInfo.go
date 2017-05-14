package definitions

type MeetingScheduleInfo struct {
	StartTime         string                           `json:"startTime,omitempty"`
	DurationInMinutes int                              `json:"durationInMinutes,omitempty"`
	TimeZone          MeetingScheduleInfo_TimezoneInfo `json:"timeZone,omitempty"`
}
