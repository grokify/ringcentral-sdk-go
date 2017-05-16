package definitions

type MeetingScheduleInfo struct {
	DurationInMinutes int                              `json:"durationInMinutes,omitempty"`
	TimeZone          MeetingScheduleInfo_TimezoneInfo `json:"timeZone,omitempty"`
	StartTime         string                           `json:"startTime,omitempty"`
}
