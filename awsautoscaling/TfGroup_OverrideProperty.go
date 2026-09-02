package awsautoscaling


// Experimental.
type TfGroup_OverrideProperty struct {
	// instance_requirements block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#instance_requirements TfGroup#instance_requirements}
	// Experimental.
	InstanceRequirements *TfGroup_InstanceRequirementsProperty `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#instance_type TfGroup#instance_type}.
	// Experimental.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// launch_template_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#launch_template_specification TfGroup#launch_template_specification}
	// Experimental.
	LaunchTemplateSpecification *TfGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationProperty `field:"optional" json:"launchTemplateSpecification" yaml:"launchTemplateSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#weighted_capacity TfGroup#weighted_capacity}.
	// Experimental.
	WeightedCapacity *string `field:"optional" json:"weightedCapacity" yaml:"weightedCapacity"`
}

