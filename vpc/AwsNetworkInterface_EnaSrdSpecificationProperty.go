package vpc


// Experimental.
type AwsNetworkInterface_EnaSrdSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#ena_srd_enabled AwsNetworkInterface#ena_srd_enabled}.
	// Experimental.
	EnaSrdEnabled interface{} `field:"optional" json:"enaSrdEnabled" yaml:"enaSrdEnabled"`
	// ena_srd_udp_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#ena_srd_udp_specification AwsNetworkInterface#ena_srd_udp_specification}
	// Experimental.
	EnaSrdUdpSpecification *AwsNetworkInterface_EnaSrdUdpSpecificationProperty `field:"optional" json:"enaSrdUdpSpecification" yaml:"enaSrdUdpSpecification"`
}

