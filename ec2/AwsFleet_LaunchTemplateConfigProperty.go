package ec2


// Experimental.
type AwsFleet_LaunchTemplateConfigProperty struct {
	// launch_template_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#launch_template_specification AwsFleet#launch_template_specification}
	// Experimental.
	LaunchTemplateSpecification *AwsFleet_LaunchTemplateSpecificationProperty `field:"optional" json:"launchTemplateSpecification" yaml:"launchTemplateSpecification"`
	// override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#override AwsFleet#override}
	// Experimental.
	Override interface{} `field:"optional" json:"override" yaml:"override"`
}

