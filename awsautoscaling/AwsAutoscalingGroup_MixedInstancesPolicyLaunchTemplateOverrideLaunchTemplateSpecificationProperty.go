package awsautoscaling


// Experimental.
type AwsAutoscalingGroup_MixedInstancesPolicyLaunchTemplateOverrideLaunchTemplateSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#launch_template_id AwsAutoscalingGroup#launch_template_id}.
	// Experimental.
	LaunchTemplateId *string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#launch_template_name AwsAutoscalingGroup#launch_template_name}.
	// Experimental.
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#version AwsAutoscalingGroup#version}.
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

