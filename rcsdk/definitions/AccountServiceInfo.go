package definitions

type AccountServiceInfo struct {
	ServiceFeatures `json:"serviceFeatures,omitempty"`
	Limits          AccountLimits   `json:"limits,omitempty"`
	Uri             string          `json:"uri,omitempty"`
	ServicePlanName string          `json:"servicePlanName,omitempty"`
	Brand           BrandInfo       `json:"brand,omitempty"`
	ServicePlan     ServicePlanInfo `json:"servicePlan,omitempty"`
	BillingPlan     BillingPlanInfo `json:"billingPlan,omitempty"`
}
