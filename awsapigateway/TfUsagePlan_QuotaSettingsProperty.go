package awsapigateway


// Experimental.
type TfUsagePlan_QuotaSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_usage_plan#limit TfUsagePlan#limit}.
	// Experimental.
	Limit *float64 `field:"required" json:"limit" yaml:"limit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_usage_plan#period TfUsagePlan#period}.
	// Experimental.
	Period *string `field:"required" json:"period" yaml:"period"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_usage_plan#offset TfUsagePlan#offset}.
	// Experimental.
	Offset *float64 `field:"optional" json:"offset" yaml:"offset"`
}

