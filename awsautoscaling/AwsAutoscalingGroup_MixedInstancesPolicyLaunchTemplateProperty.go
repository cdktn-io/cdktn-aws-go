package awsautoscaling


// Experimental.
type AwsAutoscalingGroup_MixedInstancesPolicyLaunchTemplateProperty struct {
	// launch_template_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#launch_template_specification AwsAutoscalingGroup#launch_template_specification}
	// Experimental.
	LaunchTemplateSpecification *AwsAutoscalingGroup_MixedInstancesPolicyLaunchTemplateLaunchTemplateSpecificationProperty `field:"required" json:"launchTemplateSpecification" yaml:"launchTemplateSpecification"`
	// override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#override AwsAutoscalingGroup#override}
	// Experimental.
	Override interface{} `field:"optional" json:"override" yaml:"override"`
}

