package apigateway


// Experimental.
type AwsUsagePlan_ThrottleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_usage_plan#path AwsUsagePlan#path}.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_usage_plan#burst_limit AwsUsagePlan#burst_limit}.
	// Experimental.
	BurstLimit *float64 `field:"optional" json:"burstLimit" yaml:"burstLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_usage_plan#rate_limit AwsUsagePlan#rate_limit}.
	// Experimental.
	RateLimit *float64 `field:"optional" json:"rateLimit" yaml:"rateLimit"`
}

