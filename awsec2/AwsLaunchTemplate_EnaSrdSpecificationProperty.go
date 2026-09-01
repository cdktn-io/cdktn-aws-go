package awsec2


// Experimental.
type AwsLaunchTemplate_EnaSrdSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#ena_srd_enabled AwsLaunchTemplate#ena_srd_enabled}.
	// Experimental.
	EnaSrdEnabled interface{} `field:"optional" json:"enaSrdEnabled" yaml:"enaSrdEnabled"`
	// ena_srd_udp_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#ena_srd_udp_specification AwsLaunchTemplate#ena_srd_udp_specification}
	// Experimental.
	EnaSrdUdpSpecification *AwsLaunchTemplate_EnaSrdUdpSpecificationProperty `field:"optional" json:"enaSrdUdpSpecification" yaml:"enaSrdUdpSpecification"`
}

