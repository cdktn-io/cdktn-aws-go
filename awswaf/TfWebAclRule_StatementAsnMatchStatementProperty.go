package awswaf


// Experimental.
type TfWebAclRule_StatementAsnMatchStatementProperty struct {
	// List of ASN numbers (0-4294967295).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#asn_list TfWebAclRule#asn_list}
	// Experimental.
	AsnList *[]*float64 `field:"required" json:"asnList" yaml:"asnList"`
	// forwarded_ip_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#forwarded_ip_config TfWebAclRule#forwarded_ip_config}
	// Experimental.
	ForwardedIpConfig interface{} `field:"optional" json:"forwardedIpConfig" yaml:"forwardedIpConfig"`
}

