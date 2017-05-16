package definitions

type BillingPlanInfo struct {
	Duration     string `json:"duration,omitempty"`
	DurationUnit string `json:"durationUnit,omitempty"`
	Id           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Type         string `json:"type,omitempty"`
}
