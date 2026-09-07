package autoscaling


// Experimental.
type AwsGroup_OverrideProperty struct {
	// instance_requirements block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#instance_requirements AwsGroup#instance_requirements}
	// Experimental.
	InstanceRequirements *AwsGroup_InstanceRequirementsProperty `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#instance_type AwsGroup#instance_type}.
	// Experimental.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// launch_template_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#launch_template_specification AwsGroup#launch_template_specification}
	// Experimental.
	LaunchTemplateSpecification *AwsGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationProperty `field:"optional" json:"launchTemplateSpecification" yaml:"launchTemplateSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#weighted_capacity AwsGroup#weighted_capacity}.
	// Experimental.
	WeightedCapacity *string `field:"optional" json:"weightedCapacity" yaml:"weightedCapacity"`
}

