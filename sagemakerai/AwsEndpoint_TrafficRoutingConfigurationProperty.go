package sagemakerai


// Experimental.
type AwsEndpoint_TrafficRoutingConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint#type AwsEndpoint#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint#wait_interval_in_seconds AwsEndpoint#wait_interval_in_seconds}.
	// Experimental.
	WaitIntervalInSeconds *float64 `field:"required" json:"waitIntervalInSeconds" yaml:"waitIntervalInSeconds"`
	// canary_size block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint#canary_size AwsEndpoint#canary_size}
	// Experimental.
	CanarySize *AwsEndpoint_CanarySizeProperty `field:"optional" json:"canarySize" yaml:"canarySize"`
	// linear_step_size block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint#linear_step_size AwsEndpoint#linear_step_size}
	// Experimental.
	LinearStepSize *AwsEndpoint_LinearStepSizeProperty `field:"optional" json:"linearStepSize" yaml:"linearStepSize"`
}

