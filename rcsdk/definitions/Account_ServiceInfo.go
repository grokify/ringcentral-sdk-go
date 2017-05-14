package definitions

type Account_ServiceInfo struct {
	Uri               string                `json:"uri,omitempty"`
	BillingPlan       BillingPlanInfo       `json:"billingPlan,omitempty"`
	Brand             BrandInfo             `json:"brand,omitempty"`
	ServicePlan       ServicePlanInfo       `json:"servicePlan,omitempty"`
	TargetServicePlan TargetServicePlanInfo `json:"targetServicePlan,omitempty"`
}
