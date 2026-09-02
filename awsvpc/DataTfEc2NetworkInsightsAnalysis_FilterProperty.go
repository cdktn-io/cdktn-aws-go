package awsvpc


// Experimental.
type DataTfEc2NetworkInsightsAnalysis_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_network_insights_analysis#name DataTfEc2NetworkInsightsAnalysis#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_network_insights_analysis#values DataTfEc2NetworkInsightsAnalysis#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

