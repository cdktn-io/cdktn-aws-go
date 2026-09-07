package codedeploy


// Experimental.
type AwsDeploymentGroup_TriggerConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#trigger_events AwsDeploymentGroup#trigger_events}.
	// Experimental.
	TriggerEvents *[]*string `field:"required" json:"triggerEvents" yaml:"triggerEvents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#trigger_name AwsDeploymentGroup#trigger_name}.
	// Experimental.
	TriggerName *string `field:"required" json:"triggerName" yaml:"triggerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#trigger_target_arn AwsDeploymentGroup#trigger_target_arn}.
	// Experimental.
	TriggerTargetArn *string `field:"required" json:"triggerTargetArn" yaml:"triggerTargetArn"`
}

