package awsec2


// Experimental.
type TfSpotFleetRequest_LaunchTemplateConfigProperty struct {
	// launch_template_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#launch_template_specification TfSpotFleetRequest#launch_template_specification}
	// Experimental.
	LaunchTemplateSpecification *TfSpotFleetRequest_LaunchTemplateSpecificationProperty `field:"required" json:"launchTemplateSpecification" yaml:"launchTemplateSpecification"`
	// overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#overrides TfSpotFleetRequest#overrides}
	// Experimental.
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
}

