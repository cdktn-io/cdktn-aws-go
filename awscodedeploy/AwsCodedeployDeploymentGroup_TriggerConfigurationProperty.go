package awscodedeploy


// Experimental.
type AwsCodedeployDeploymentGroup_TriggerConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#trigger_events AwsCodedeployDeploymentGroup#trigger_events}.
	// Experimental.
	TriggerEvents *[]*string `field:"required" json:"triggerEvents" yaml:"triggerEvents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#trigger_name AwsCodedeployDeploymentGroup#trigger_name}.
	// Experimental.
	TriggerName *string `field:"required" json:"triggerName" yaml:"triggerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#trigger_target_arn AwsCodedeployDeploymentGroup#trigger_target_arn}.
	// Experimental.
	TriggerTargetArn *string `field:"required" json:"triggerTargetArn" yaml:"triggerTargetArn"`
}

