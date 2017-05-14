package definitions

type Account_ServiceInfo struct {
	BillingPlan       BillingPlanInfo       `json:"billingPlan,omitempty"`
	Brand             BrandInfo             `json:"brand,omitempty"`
	ServicePlan       ServicePlanInfo       `json:"servicePlan,omitempty"`
	TargetServicePlan TargetServicePlanInfo `json:"targetServicePlan,omitempty"`
	Uri               string                `json:"uri,omitempty"`
}
