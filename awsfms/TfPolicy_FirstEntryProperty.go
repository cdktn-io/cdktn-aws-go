package awsfms


// Experimental.
type TfPolicy_FirstEntryProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#egress TfPolicy#egress}.
	// Experimental.
	Egress interface{} `field:"required" json:"egress" yaml:"egress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#protocol TfPolicy#protocol}.
	// Experimental.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#rule_action TfPolicy#rule_action}.
	// Experimental.
	RuleAction *string `field:"required" json:"ruleAction" yaml:"ruleAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#cidr_block TfPolicy#cidr_block}.
	// Experimental.
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
	// icmp_type_code block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#icmp_type_code TfPolicy#icmp_type_code}
	// Experimental.
	IcmpTypeCode interface{} `field:"optional" json:"icmpTypeCode" yaml:"icmpTypeCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#ipv6_cidr_block TfPolicy#ipv6_cidr_block}.
	// Experimental.
	Ipv6CidrBlock *string `field:"optional" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
	// port_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fms_policy#port_range TfPolicy#port_range}
	// Experimental.
	PortRange interface{} `field:"optional" json:"portRange" yaml:"portRange"`
}

