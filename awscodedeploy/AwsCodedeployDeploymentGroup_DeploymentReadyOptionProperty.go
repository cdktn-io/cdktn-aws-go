package awscodedeploy


// Experimental.
type AwsCodedeployDeploymentGroup_DeploymentReadyOptionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#action_on_timeout AwsCodedeployDeploymentGroup#action_on_timeout}.
	// Experimental.
	ActionOnTimeout *string `field:"optional" json:"actionOnTimeout" yaml:"actionOnTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#wait_time_in_minutes AwsCodedeployDeploymentGroup#wait_time_in_minutes}.
	// Experimental.
	WaitTimeInMinutes *float64 `field:"optional" json:"waitTimeInMinutes" yaml:"waitTimeInMinutes"`
}

