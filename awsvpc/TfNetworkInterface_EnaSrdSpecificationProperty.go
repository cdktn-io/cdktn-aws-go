package awsvpc


// Experimental.
type TfNetworkInterface_EnaSrdSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#ena_srd_enabled TfNetworkInterface#ena_srd_enabled}.
	// Experimental.
	EnaSrdEnabled interface{} `field:"optional" json:"enaSrdEnabled" yaml:"enaSrdEnabled"`
	// ena_srd_udp_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_interface#ena_srd_udp_specification TfNetworkInterface#ena_srd_udp_specification}
	// Experimental.
	EnaSrdUdpSpecification *TfNetworkInterface_EnaSrdUdpSpecificationProperty `field:"optional" json:"enaSrdUdpSpecification" yaml:"enaSrdUdpSpecification"`
}

