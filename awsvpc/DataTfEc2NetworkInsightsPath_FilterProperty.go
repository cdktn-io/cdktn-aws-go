package awsvpc


// Experimental.
type DataTfEc2NetworkInsightsPath_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_network_insights_path#name DataTfEc2NetworkInsightsPath#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_network_insights_path#values DataTfEc2NetworkInsightsPath#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

