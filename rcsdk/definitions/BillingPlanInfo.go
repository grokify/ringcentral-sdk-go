package definitions

type BillingPlanInfo struct {
	Id           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	DurationUnit string `json:"durationUnit,omitempty"`
	Duration     string `json:"duration,omitempty"`
	Type         string `json:"type,omitempty"`
}
