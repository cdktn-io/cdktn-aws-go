package awsapigateway


// Experimental.
type TfUsagePlan_ThrottleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_usage_plan#path TfUsagePlan#path}.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_usage_plan#burst_limit TfUsagePlan#burst_limit}.
	// Experimental.
	BurstLimit *float64 `field:"optional" json:"burstLimit" yaml:"burstLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_usage_plan#rate_limit TfUsagePlan#rate_limit}.
	// Experimental.
	RateLimit *float64 `field:"optional" json:"rateLimit" yaml:"rateLimit"`
}

