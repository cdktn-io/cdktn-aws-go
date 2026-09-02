package awsautoscaling


// Experimental.
type TfGroup_MixedInstancesPolicyProperty struct {
	// launch_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#launch_template TfGroup#launch_template}
	// Experimental.
	LaunchTemplate *TfGroup_MixedInstancesPolicyLaunchTemplateProperty `field:"required" json:"launchTemplate" yaml:"launchTemplate"`
	// instances_distribution block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#instances_distribution TfGroup#instances_distribution}
	// Experimental.
	InstancesDistribution *TfGroup_InstancesDistributionProperty `field:"optional" json:"instancesDistribution" yaml:"instancesDistribution"`
}

