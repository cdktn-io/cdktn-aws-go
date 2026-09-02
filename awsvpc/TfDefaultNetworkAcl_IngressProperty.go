package awsvpc


// Experimental.
type TfDefaultNetworkAcl_IngressProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_network_acl#action TfDefaultNetworkAcl#action}.
	// Experimental.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_network_acl#from_port TfDefaultNetworkAcl#from_port}.
	// Experimental.
	FromPort *float64 `field:"required" json:"fromPort" yaml:"fromPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_network_acl#protocol TfDefaultNetworkAcl#protocol}.
	// Experimental.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_network_acl#rule_no TfDefaultNetworkAcl#rule_no}.
	// Experimental.
	RuleNo *float64 `field:"required" json:"ruleNo" yaml:"ruleNo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_network_acl#to_port TfDefaultNetworkAcl#to_port}.
	// Experimental.
	ToPort *float64 `field:"required" json:"toPort" yaml:"toPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_network_acl#cidr_block TfDefaultNetworkAcl#cidr_block}.
	// Experimental.
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_network_acl#icmp_code TfDefaultNetworkAcl#icmp_code}.
	// Experimental.
	IcmpCode *float64 `field:"optional" json:"icmpCode" yaml:"icmpCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_network_acl#icmp_type TfDefaultNetworkAcl#icmp_type}.
	// Experimental.
	IcmpType *float64 `field:"optional" json:"icmpType" yaml:"icmpType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/default_network_acl#ipv6_cidr_block TfDefaultNetworkAcl#ipv6_cidr_block}.
	// Experimental.
	Ipv6CidrBlock *string `field:"optional" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
}

