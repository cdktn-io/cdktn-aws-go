package awsec2


// Experimental.
type TfFleet_LaunchTemplateConfigProperty struct {
	// launch_template_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#launch_template_specification TfFleet#launch_template_specification}
	// Experimental.
	LaunchTemplateSpecification *TfFleet_LaunchTemplateSpecificationProperty `field:"optional" json:"launchTemplateSpecification" yaml:"launchTemplateSpecification"`
	// override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#override TfFleet#override}
	// Experimental.
	Override interface{} `field:"optional" json:"override" yaml:"override"`
}

