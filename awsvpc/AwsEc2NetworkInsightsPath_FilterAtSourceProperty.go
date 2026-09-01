package awsvpc


// Experimental.
type AwsEc2NetworkInsightsPath_FilterAtSourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_network_insights_path#destination_address AwsEc2NetworkInsightsPath#destination_address}.
	// Experimental.
	DestinationAddress *string `field:"optional" json:"destinationAddress" yaml:"destinationAddress"`
	// destination_port_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_network_insights_path#destination_port_range AwsEc2NetworkInsightsPath#destination_port_range}
	// Experimental.
	DestinationPortRange *AwsEc2NetworkInsightsPath_FilterAtSourceDestinationPortRangeProperty `field:"optional" json:"destinationPortRange" yaml:"destinationPortRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_network_insights_path#source_address AwsEc2NetworkInsightsPath#source_address}.
	// Experimental.
	SourceAddress *string `field:"optional" json:"sourceAddress" yaml:"sourceAddress"`
	// source_port_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_network_insights_path#source_port_range AwsEc2NetworkInsightsPath#source_port_range}
	// Experimental.
	SourcePortRange *AwsEc2NetworkInsightsPath_FilterAtSourceSourcePortRangeProperty `field:"optional" json:"sourcePortRange" yaml:"sourcePortRange"`
}

