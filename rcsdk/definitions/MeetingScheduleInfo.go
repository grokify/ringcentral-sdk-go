package definitions

type MeetingScheduleInfo struct {
	DurationInMinutes int                              `json:"durationInMinutes,omitempty"`
	StartTime         string                           `json:"startTime,omitempty"`
	TimeZone          MeetingScheduleInfo_TimezoneInfo `json:"timeZone,omitempty"`
}
