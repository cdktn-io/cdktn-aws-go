package codedeploy


// Experimental.
type AwsDeploymentConfig_TrafficRoutingConfigProperty struct {
	// time_based_canary block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_config#time_based_canary AwsDeploymentConfig#time_based_canary}
	// Experimental.
	TimeBasedCanary *AwsDeploymentConfig_TimeBasedCanaryProperty `field:"optional" json:"timeBasedCanary" yaml:"timeBasedCanary"`
	// time_based_linear block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_config#time_based_linear AwsDeploymentConfig#time_based_linear}
	// Experimental.
	TimeBasedLinear *AwsDeploymentConfig_TimeBasedLinearProperty `field:"optional" json:"timeBasedLinear" yaml:"timeBasedLinear"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_config#type AwsDeploymentConfig#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

