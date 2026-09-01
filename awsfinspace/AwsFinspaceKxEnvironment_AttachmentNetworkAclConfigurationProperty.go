package awsfinspace


// Experimental.
type AwsFinspaceKxEnvironment_AttachmentNetworkAclConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_environment#cidr_block AwsFinspaceKxEnvironment#cidr_block}.
	// Experimental.
	CidrBlock *string `field:"required" json:"cidrBlock" yaml:"cidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_environment#protocol AwsFinspaceKxEnvironment#protocol}.
	// Experimental.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_environment#rule_action AwsFinspaceKxEnvironment#rule_action}.
	// Experimental.
	RuleAction *string `field:"required" json:"ruleAction" yaml:"ruleAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_environment#rule_number AwsFinspaceKxEnvironment#rule_number}.
	// Experimental.
	RuleNumber *float64 `field:"required" json:"ruleNumber" yaml:"ruleNumber"`
	// icmp_type_code block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_environment#icmp_type_code AwsFinspaceKxEnvironment#icmp_type_code}
	// Experimental.
	IcmpTypeCode *AwsFinspaceKxEnvironment_IcmpTypeCodeProperty `field:"optional" json:"icmpTypeCode" yaml:"icmpTypeCode"`
	// port_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_environment#port_range AwsFinspaceKxEnvironment#port_range}
	// Experimental.
	PortRange *AwsFinspaceKxEnvironment_PortRangeProperty `field:"optional" json:"portRange" yaml:"portRange"`
}

