package codedeploy


// Experimental.
type AwsDeploymentGroup_TerminateBlueInstancesOnDeploymentSuccessProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#action AwsDeploymentGroup#action}.
	// Experimental.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#termination_wait_time_in_minutes AwsDeploymentGroup#termination_wait_time_in_minutes}.
	// Experimental.
	TerminationWaitTimeInMinutes *float64 `field:"optional" json:"terminationWaitTimeInMinutes" yaml:"terminationWaitTimeInMinutes"`
}

