package awscloudfront


// Experimental.
type TfContinuousDeploymentPolicy_SingleWeightConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_continuous_deployment_policy#weight TfContinuousDeploymentPolicy#weight}.
	// Experimental.
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
	// session_stickiness_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_continuous_deployment_policy#session_stickiness_config TfContinuousDeploymentPolicy#session_stickiness_config}
	// Experimental.
	SessionStickinessConfig interface{} `field:"optional" json:"sessionStickinessConfig" yaml:"sessionStickinessConfig"`
}

