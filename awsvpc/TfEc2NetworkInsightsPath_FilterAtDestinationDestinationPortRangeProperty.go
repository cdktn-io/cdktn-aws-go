package awsvpc


// Experimental.
type TfEc2NetworkInsightsPath_FilterAtDestinationDestinationPortRangeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_network_insights_path#from_port TfEc2NetworkInsightsPath#from_port}.
	// Experimental.
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_network_insights_path#to_port TfEc2NetworkInsightsPath#to_port}.
	// Experimental.
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

