package awsec2


// Experimental.
type AwsSpotFleetRequest_LaunchTemplateConfigProperty struct {
	// launch_template_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#launch_template_specification AwsSpotFleetRequest#launch_template_specification}
	// Experimental.
	LaunchTemplateSpecification *AwsSpotFleetRequest_LaunchTemplateSpecificationProperty `field:"required" json:"launchTemplateSpecification" yaml:"launchTemplateSpecification"`
	// overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#overrides AwsSpotFleetRequest#overrides}
	// Experimental.
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
}

