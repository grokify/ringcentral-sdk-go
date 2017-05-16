package definitions

type AccountServiceInfo struct {
	BillingPlan     BillingPlanInfo      `json:"billingPlan,omitempty"`
	Brand           BrandInfo            `json:"brand,omitempty"`
	Limits          AccountLimits        `json:"limits,omitempty"`
	ServiceFeatures []ServiceFeatureInfo `json:"serviceFeatures,omitempty"`
	ServicePlan     ServicePlanInfo      `json:"servicePlan,omitempty"`
	ServicePlanName string               `json:"servicePlanName,omitempty"`
	Uri             string               `json:"uri,omitempty"`
}
