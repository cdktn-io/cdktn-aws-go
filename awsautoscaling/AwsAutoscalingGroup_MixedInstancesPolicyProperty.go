package awsautoscaling


// Experimental.
type AwsAutoscalingGroup_MixedInstancesPolicyProperty struct {
	// launch_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#launch_template AwsAutoscalingGroup#launch_template}
	// Experimental.
	LaunchTemplate *AwsAutoscalingGroup_MixedInstancesPolicyLaunchTemplateProperty `field:"required" json:"launchTemplate" yaml:"launchTemplate"`
	// instances_distribution block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#instances_distribution AwsAutoscalingGroup#instances_distribution}
	// Experimental.
	InstancesDistribution *AwsAutoscalingGroup_InstancesDistributionProperty `field:"optional" json:"instancesDistribution" yaml:"instancesDistribution"`
}

