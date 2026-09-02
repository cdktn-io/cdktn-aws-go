package awswaf


// Experimental.
type TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementIpSetReferenceStatementProperty struct {
	// ARN of the IP set to reference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#arn TfWebAclRule#arn}
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// ip_set_forwarded_ip_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafv2_web_acl_rule#ip_set_forwarded_ip_config TfWebAclRule#ip_set_forwarded_ip_config}
	// Experimental.
	IpSetForwardedIpConfig interface{} `field:"optional" json:"ipSetForwardedIpConfig" yaml:"ipSetForwardedIpConfig"`
}

