package cloudfront


// Experimental.
type AwsContinuousDeploymentPolicy_SessionStickinessConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_continuous_deployment_policy#idle_ttl AwsContinuousDeploymentPolicy#idle_ttl}.
	// Experimental.
	IdleTtl *float64 `field:"required" json:"idleTtl" yaml:"idleTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_continuous_deployment_policy#maximum_ttl AwsContinuousDeploymentPolicy#maximum_ttl}.
	// Experimental.
	MaximumTtl *float64 `field:"required" json:"maximumTtl" yaml:"maximumTtl"`
}

