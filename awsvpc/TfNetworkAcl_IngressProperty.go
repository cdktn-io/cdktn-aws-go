package awsvpc


// Experimental.
type TfNetworkAcl_IngressProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_acl#action TfNetworkAcl#action}.
	// Experimental.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_acl#cidr_block TfNetworkAcl#cidr_block}.
	// Experimental.
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_acl#from_port TfNetworkAcl#from_port}.
	// Experimental.
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_acl#icmp_code TfNetworkAcl#icmp_code}.
	// Experimental.
	IcmpCode *float64 `field:"optional" json:"icmpCode" yaml:"icmpCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_acl#icmp_type TfNetworkAcl#icmp_type}.
	// Experimental.
	IcmpType *float64 `field:"optional" json:"icmpType" yaml:"icmpType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_acl#ipv6_cidr_block TfNetworkAcl#ipv6_cidr_block}.
	// Experimental.
	Ipv6CidrBlock *string `field:"optional" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_acl#protocol TfNetworkAcl#protocol}.
	// Experimental.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_acl#rule_no TfNetworkAcl#rule_no}.
	// Experimental.
	RuleNo *float64 `field:"optional" json:"ruleNo" yaml:"ruleNo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/network_acl#to_port TfNetworkAcl#to_port}.
	// Experimental.
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

