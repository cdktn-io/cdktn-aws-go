package finspace


// Experimental.
type AwsKxEnvironment_AttachmentNetworkAclConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_environment#cidr_block AwsKxEnvironment#cidr_block}.
	// Experimental.
	CidrBlock *string `field:"required" json:"cidrBlock" yaml:"cidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_environment#protocol AwsKxEnvironment#protocol}.
	// Experimental.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_environment#rule_action AwsKxEnvironment#rule_action}.
	// Experimental.
	RuleAction *string `field:"required" json:"ruleAction" yaml:"ruleAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_environment#rule_number AwsKxEnvironment#rule_number}.
	// Experimental.
	RuleNumber *float64 `field:"required" json:"ruleNumber" yaml:"ruleNumber"`
	// icmp_type_code block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_environment#icmp_type_code AwsKxEnvironment#icmp_type_code}
	// Experimental.
	IcmpTypeCode *AwsKxEnvironment_IcmpTypeCodeProperty `field:"optional" json:"icmpTypeCode" yaml:"icmpTypeCode"`
	// port_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_environment#port_range AwsKxEnvironment#port_range}
	// Experimental.
	PortRange *AwsKxEnvironment_PortRangeProperty `field:"optional" json:"portRange" yaml:"portRange"`
}

