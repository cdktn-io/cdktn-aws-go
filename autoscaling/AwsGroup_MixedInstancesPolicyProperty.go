package autoscaling


// Experimental.
type AwsGroup_MixedInstancesPolicyProperty struct {
	// launch_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#launch_template AwsGroup#launch_template}
	// Experimental.
	LaunchTemplate *AwsGroup_MixedInstancesPolicyLaunchTemplateProperty `field:"required" json:"launchTemplate" yaml:"launchTemplate"`
	// instances_distribution block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#instances_distribution AwsGroup#instances_distribution}
	// Experimental.
	InstancesDistribution *AwsGroup_InstancesDistributionProperty `field:"optional" json:"instancesDistribution" yaml:"instancesDistribution"`
}

