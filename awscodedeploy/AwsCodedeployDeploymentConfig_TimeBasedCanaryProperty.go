package awscodedeploy


// Experimental.
type AwsCodedeployDeploymentConfig_TimeBasedCanaryProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_config#interval AwsCodedeployDeploymentConfig#interval}.
	// Experimental.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_config#percentage AwsCodedeployDeploymentConfig#percentage}.
	// Experimental.
	Percentage *float64 `field:"optional" json:"percentage" yaml:"percentage"`
}

