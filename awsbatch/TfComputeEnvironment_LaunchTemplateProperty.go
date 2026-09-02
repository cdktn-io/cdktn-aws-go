package awsbatch


// Experimental.
type TfComputeEnvironment_LaunchTemplateProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_compute_environment#launch_template_id TfComputeEnvironment#launch_template_id}.
	// Experimental.
	LaunchTemplateId *string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_compute_environment#launch_template_name TfComputeEnvironment#launch_template_name}.
	// Experimental.
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_compute_environment#version TfComputeEnvironment#version}.
	// Experimental.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

