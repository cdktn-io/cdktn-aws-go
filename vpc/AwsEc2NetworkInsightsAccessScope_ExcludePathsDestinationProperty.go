package vpc


// Experimental.
type AwsEc2NetworkInsightsAccessScope_ExcludePathsDestinationProperty struct {
	// packet_header_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_network_insights_access_scope#packet_header_statement AwsEc2NetworkInsightsAccessScope#packet_header_statement}
	// Experimental.
	PacketHeaderStatement interface{} `field:"optional" json:"packetHeaderStatement" yaml:"packetHeaderStatement"`
	// resource_statement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_network_insights_access_scope#resource_statement AwsEc2NetworkInsightsAccessScope#resource_statement}
	// Experimental.
	ResourceStatement interface{} `field:"optional" json:"resourceStatement" yaml:"resourceStatement"`
}

