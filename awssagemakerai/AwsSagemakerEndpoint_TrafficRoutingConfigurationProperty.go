package awssagemakerai


// Experimental.
type AwsSagemakerEndpoint_TrafficRoutingConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint#type AwsSagemakerEndpoint#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint#wait_interval_in_seconds AwsSagemakerEndpoint#wait_interval_in_seconds}.
	// Experimental.
	WaitIntervalInSeconds *float64 `field:"required" json:"waitIntervalInSeconds" yaml:"waitIntervalInSeconds"`
	// canary_size block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint#canary_size AwsSagemakerEndpoint#canary_size}
	// Experimental.
	CanarySize *AwsSagemakerEndpoint_CanarySizeProperty `field:"optional" json:"canarySize" yaml:"canarySize"`
	// linear_step_size block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint#linear_step_size AwsSagemakerEndpoint#linear_step_size}
	// Experimental.
	LinearStepSize *AwsSagemakerEndpoint_LinearStepSizeProperty `field:"optional" json:"linearStepSize" yaml:"linearStepSize"`
}

