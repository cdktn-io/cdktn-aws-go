package vpc


// Experimental.
type AwsEc2NetworkInsightsAccessScope_ExcludePathsDestinationPacketHeaderStatementProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_network_insights_access_scope#destination_addresses AwsEc2NetworkInsightsAccessScope#destination_addresses}.
	// Experimental.
	DestinationAddresses *[]*string `field:"optional" json:"destinationAddresses" yaml:"destinationAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_network_insights_access_scope#destination_ports AwsEc2NetworkInsightsAccessScope#destination_ports}.
	// Experimental.
	DestinationPorts *[]*string `field:"optional" json:"destinationPorts" yaml:"destinationPorts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_network_insights_access_scope#destination_prefix_lists AwsEc2NetworkInsightsAccessScope#destination_prefix_lists}.
	// Experimental.
	DestinationPrefixLists *[]*string `field:"optional" json:"destinationPrefixLists" yaml:"destinationPrefixLists"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_network_insights_access_scope#protocols AwsEc2NetworkInsightsAccessScope#protocols}.
	// Experimental.
	Protocols *[]*string `field:"optional" json:"protocols" yaml:"protocols"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_network_insights_access_scope#source_addresses AwsEc2NetworkInsightsAccessScope#source_addresses}.
	// Experimental.
	SourceAddresses *[]*string `field:"optional" json:"sourceAddresses" yaml:"sourceAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_network_insights_access_scope#source_ports AwsEc2NetworkInsightsAccessScope#source_ports}.
	// Experimental.
	SourcePorts *[]*string `field:"optional" json:"sourcePorts" yaml:"sourcePorts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_network_insights_access_scope#source_prefix_lists AwsEc2NetworkInsightsAccessScope#source_prefix_lists}.
	// Experimental.
	SourcePrefixLists *[]*string `field:"optional" json:"sourcePrefixLists" yaml:"sourcePrefixLists"`
}

