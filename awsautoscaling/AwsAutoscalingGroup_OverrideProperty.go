package awsautoscaling


// Experimental.
type AwsAutoscalingGroup_OverrideProperty struct {
	// instance_requirements block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#instance_requirements AwsAutoscalingGroup#instance_requirements}
	// Experimental.
	InstanceRequirements *AwsAutoscalingGroup_InstanceRequirementsProperty `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#instance_type AwsAutoscalingGroup#instance_type}.
	// Experimental.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// launch_template_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#launch_template_specification AwsAutoscalingGroup#launch_template_specification}
	// Experimental.
	LaunchTemplateSpecification *AwsAutoscalingGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationProperty `field:"optional" json:"launchTemplateSpecification" yaml:"launchTemplateSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#weighted_capacity AwsAutoscalingGroup#weighted_capacity}.
	// Experimental.
	WeightedCapacity *string `field:"optional" json:"weightedCapacity" yaml:"weightedCapacity"`
}

